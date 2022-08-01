package app

import (
	"capi/domain"
	"capi/errs"
	"capi/logger"
	"capi/service"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}

	for _, envKey := range envProps {
		if os.Getenv(envKey) == "" {
			logger.Fatal(fmt.Sprintf("environment variable %s not defined. terminating application...", envKey))
		}
	}

	logger.Info("environment variables loaded...")

}

func Start() {

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("error loading .env file")
	}
	logger.Info("load environment variables...")

	sanityCheck()

	dbClient := getClientDB()

	// * wiring
	// * setup repository
	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDB(dbClient)
	authRepository := domain.NewAuthRepository(dbClient)

	// * setup service
	customerService := service.NewCustomerService(customerRepositoryDB)
	accountService := service.NewAccountService(accountRepositoryDB)
	authService := service.NewAuthService(authRepository)

	// * setup handler
	ch := CustomerHandlers{customerService}
	ah := AccountHandler{accountService}
	authHandler := AuthHandler{authService}

	// * create ServeMux
	mux := mux.NewRouter()

	mux.Use(loggingMiddleware)

	// * defining routes
	authR := mux.PathPrefix("/auth").Subrouter()
	authR.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)

	customerR := mux.PathPrefix("/customers").Subrouter()
	customerR.HandleFunc("/{customer_id:[0-9]+}", ch.getCustomerByID).Methods(http.MethodGet)
	customerR.HandleFunc("/{customer_id:[0-9]+}/accounts/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)
	customerR.Use(authMiddleware)

	adminR := mux.PathPrefix("/customers").Subrouter()
	adminR.HandleFunc("", ch.getAllCustomers).Methods(http.MethodGet)
	adminR.HandleFunc("/{customer_id:[0-9]+}/accounts", ah.NewAccount).Methods(http.MethodPost)
	adminR.Use(authMiddleware)
	adminR.Use(isAdminMiddleware)

	// * starting the server

	serverAddr := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")

	logger.Info(fmt.Sprintf("start server on %s:%s...", serverAddr, serverPort))
	http.ListenAndServe(fmt.Sprintf("%s:%s", serverAddr, serverPort), mux)
}

func getClientDB() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("success connect to database...")

	return db
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		timer := time.Now()

		next.ServeHTTP(w, r)

		logger.Info(fmt.Sprintf("%v %v %v", r.Method, r.URL, time.Since(timer)))
	})
}

type key int

const (
	userInfo key = iota
	// ...
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			appErr := errs.NewBadRequestError("invalid token")
			writeResponse(w, appErr.Code, appErr.AsMessage())
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.ParseWithClaims(tokenString, &domain.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("signing method invalid")
			}

			return []byte("hmacSampleSecret"), nil
		})

		if err != nil {
			appErr := errs.NewBadRequestError(err.Error())
			writeResponse(w, appErr.Code, appErr.AsMessage())
			return
		}

		claims, ok := token.Claims.(*domain.AccessTokenClaims)
		if !ok || !token.Valid {
			appErr := errs.NewBadRequestError("invalid token")
			writeResponse(w, appErr.Code, appErr.AsMessage())
			return
		}

		if claims.IsUserRole() {
			vars := mux.Vars(r)
			logger.Info(fmt.Sprintf("vars: %v", vars))

			if !claims.IsRequestVerifiedWithTokenClaims(vars) {
				appErr := errs.NewAuthenticationError("request not verified with the token claims")
				writeResponse(w, appErr.Code, appErr.AsMessage())
				return
			}
		}

		ctx := context.WithValue(r.Context(), userInfo, claims)
		r = r.WithContext(ctx)

		vars2 := mux.Vars(r)
		logger.Info(fmt.Sprintf("vars2: %v", vars2))

		next.ServeHTTP(w, r)

	})
}

func isAdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context().Value(userInfo).(*domain.AccessTokenClaims)

		if ctx.Role != "admin" {
			appErr := errs.NewForbiddenError("don't have enough permission to get the resources")
			writeResponse(w, appErr.Code, appErr.AsMessage())
			return
		}

		next.ServeHTTP(w, r)

	})

}
