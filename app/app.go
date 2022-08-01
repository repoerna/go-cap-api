package app

import (
	"capi/domain"
	"capi/logger"
	"capi/service"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func sanityCheck(){
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWORD",
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
	
	// * setup handle
	customerService := service.NewCustomerService(customerRepositoryDB)
	accountService := service.NewAccountService(accountRepositoryDB)
	
	// * wiring
	ch := CustomerHandler{customerService}
	ah := AccountHandler{accountService}

	// * create ServeRoute
	mux := mux.NewRouter()

	mux.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")

	mux.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerByID).Methods("GET")

	
	mux.HandleFunc("/customers/{customer_id:[0-9]+}/accounts", ah.NewAccount).Methods(http.MethodPost)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}/accounts/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)

	// * starting the server
	fmt.Println("starting the server localhost:9000")
	http.ListenAndServe(":9000", mux)
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