package app

import (
	"capi/domain"
	"capi/errs"
	"capi/logger"
	"capi/service"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func sanityCheck(){
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
			logger.Fatal(fmt.Sprintf("environtment variabel %s not defined. terminating application..", envKey))
		}
	}

	logger.Info("environtment variabel loaded...")
}

func Start() {
	
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("error loading .env file...")
	}

	logger.Info("load environment variabel...")

	sanityCheck()

	dbClient := getClientDB()
	
	// * wiring
	// * setup repository
	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDB(dbClient)
	authRepositoryDB := domain.NewAuthRepositoryDB(dbClient)
	
	// * setup handle
	customerService := service.NewCustomerService(customerRepositoryDB)
	accountService := service.NewAccountService(accountRepositoryDB)
	authService := service.NewAuthService(authRepositoryDB)
	
	// * wiring
	ch := CustomerHandler{customerService}
	ah := AccountHandler{accountService}
	authH := AuthHandler{authService}

	// * create ServeRoute
	mux := mux.NewRouter()

	authR := mux.PathPrefix("/auth").Subrouter()
	authR.HandleFunc("/login", authH.Login).Methods(http.MethodPost)

	authR.Use(loggingMiddleware)

	mux.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")

	mux.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerByID).Methods("GET")

	
	mux.HandleFunc("/customers/{customer_id:[0-9]+}/accounts", ah.NewAccount).Methods(http.MethodPost)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}/accounts/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)

	mux.Use(authMiddleware)
	// * starting the server
	serverAddr := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")

	logger.Info(fmt.Sprintf("Start Server on %s:%s...", serverAddr, serverPort))
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
	logger.Info("Success Connect to Database...")

	return db
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		timer := time.Now()
		next.ServeHTTP(w, r)
		logger.Info(fmt.Sprintf("%v %v %v", r.Method, r.URL, time.Since(timer)))
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		token  := r.Header.Get("Authorization")

		// check token validation
		if !strings.Contains(token, "Bearer") {
			logger.Error("Token Invalid")
			errApp := errs.NewForbiddenError("Invalid Token")
			writeResponse(w, errApp.Code, errApp.AsMessage())
			return
		}

		// spilt token -> ambil tokennya buang "Bearer" nya
		tokenArr := strings.Split(token, " ") // me-return array
		
		tokenString := ""

		if len(tokenArr) == 2 {
			tokenString = tokenArr[1]
		}

		// parsing token, err != jwt.parse(
		logger.Info("THE TOKEN IS: " + tokenString)
		signedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {return []byte("rahasia"), nil})

		fmt.Println(err)
		fmt.Println("TEST SIGNIN TOKEN: ", signedToken)
		if signedToken.Valid {
			fmt.Println("You look nice today")
		} else if errors.Is(err, jwt.ErrTokenMalformed) {
			fmt.Println("That's not even a token")
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}

		logger.Info(token)

		next.ServeHTTP(w, r)
	})
}