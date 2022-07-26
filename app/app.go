package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// * create ServeRoute
	mux := mux.NewRouter()

	// * defining routes
	mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
	mux.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	// mux.HandleFunc("/customers", getAllCustomers)
	mux.HandleFunc("/customer/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	//DELETE
	mux.HandleFunc("/customer/{customer_id:[0-9]+}", deleteCustomer).Methods("DELETE")
	// Post
	mux.HandleFunc("/customers", addCustomer).Methods("post")
	// PUT
	mux.HandleFunc("/customer/{customer_id:[0-9]+}", updateCustomer).Methods("PUT")

	// * starting the server
	fmt.Println("starting the server localhost:9000")
	http.ListenAndServe(":9000", mux)
}