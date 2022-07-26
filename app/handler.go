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
	ID int `json:"id" xml:"id"`
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipCode"`
}

var customers []Customer = []Customer{
	{1, "Bruyne", "Banten", "15540"},
	{15, "Brown", "Surabaya", "15560"},
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Celerates!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Get Custoem Endpoint!")

	// w.Header().Add("Content Type", "application/json")
	// json.NewEncoder(w).Encode(customers)
	// xml.NewEncoder(w).Encode(customers)

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	// get route variable
	vars := mux.Vars(r)

	customerId := vars["customer_id"]

	//convert string to int
	id, err := strconv.Atoi(customerId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Print("\n", w, "Invalid Customer id\n")
	}

	// Searching customer data
	var cust Customer

	for _, data := range customers {
		if data.ID == id {
			cust = data
		}
	}

	if cust.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Customer id Not Found\n")
		return
	}

	// return customer data
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cust)
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	var cust Customer
	json.NewDecoder(r.Body).Decode(&cust)

	// get last id
	nextID := getNextID()
	cust.ID = nextID

	// * save data to array
	customers = append(customers, cust)

	w.WriteHeader(http.StatusCreated)

	fmt.Fprintln(w, "Customer Successfully Added!")

	json.NewEncoder(w).Encode(customers)
}

func getNextID() int {
	// lastIndex := len(customers) - 1
	// lastCustomer := 
	cust := customers[len(customers)-1]
	return cust.ID + 1
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	customerId := vars["customer_id"]

	//convert string to int
	id, err := strconv.Atoi(customerId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Print("\n", w, "Invalid Customer id\n")
		return
	}

	for index, cust := range customers {
		if cust.ID == id {
			customers = append(customers[:index], customers[index+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Customer Successfully Deleted!")

	json.NewEncoder(w).Encode(customers)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	vars := mux.Vars(r)
	
	customerId := vars["customer_id"]
	//convert string to int
	id, err := strconv.Atoi(customerId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Print("\n", w, "Invalid Customer id\n")
		return
	}
	
	
	// var cust Customer
	for index, data := range customers {

		if data.ID == id {
			var newCustomer Customer
			json.NewDecoder(r.Body).Decode(&newCustomer)

			customers[index].Name = newCustomer.Name
			customers[index].City = newCustomer.City
			customers[index].ZipCode = newCustomer.ZipCode

			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Customer Successfully Updated!")

			json.NewEncoder(w).Encode(customers)
		}
	}

}