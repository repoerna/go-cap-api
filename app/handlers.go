package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

var customers []Customer = []Customer{
	{"User 1", "Jakarta", "12345"},
	{"User 2", "Surabaya", "67890"},
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Celerates!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	customerId := vars["customer_id"]

	id, err := strconv.Atoi(customerId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if len(customers) < id || id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("customer id not found")
		return
	}

	customerData := customers[id-1]

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customerData)
}
