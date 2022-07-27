package app

import (
	"capi/domain"
	"capi/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// wiring
	ch := CustomerHendler{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	//create mux
	mux := mux.NewRouter()

	// * defining routes
	// mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
	mux.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	// mux.HandleFunc("/customers", addCustomer).Methods(http.MethodPost)
	
	
	// mux.HandleFunc("/customers/{customer_id:[0-9]+}", updateCustomer).Methods(http.MethodPut)
	// mux.HandleFunc("/customers/{customer_id:[0-9]+}", deleteCustomer).Methods(http.MethodDelete)
	// mux.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// * starting the server
	http.ListenAndServe(":8080", mux)
}