package app

import (
	"capi/service"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

// type Customer struct {
// 	Name    string `json:"name" xml:"name"`
// 	City    string `json:"city" xml:"city"`
// 	Zipcode string `json:"zip_code" xml:"zipcode"`
// }

// var customers []Customer = []Customer{
// 	{"User 1", "Jakarta", "12345"},
// 	{"User 2", "Surabaya", "67890"},
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello Celerates!")
// }

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

func (ch *CustomerHandlers) getCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	customerID := vars["customer_id"]

	// id, err := strconv.Atoi(customerId)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	// if len(customers) < id || id == 0 {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	json.NewEncoder(w).Encode("customer id not found")
	// 	return
	// }

	customer, err := ch.service.GetCustomerByID(customerID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
