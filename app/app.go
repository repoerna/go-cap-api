package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// * create multiplexer
	router := mux.NewRouter()

	// * defining routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)

	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// * starting the server
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "post request received")
}
