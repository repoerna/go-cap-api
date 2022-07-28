package app

import (
	"capi/domain"
	"capi/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// * wiring
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// * create ServeRoute
	mux := mux.NewRouter()

	mux.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")

	mux.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerByID).Methods("GET")

	// * starting the server
	fmt.Println("starting the server localhost:9000")
	http.ListenAndServe(":9000", mux)
}