package app

import (
	"capi/service"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Get Custoem Endpoint!")

	// w.Header().Add("Content Type", "application/json")
	// json.NewEncoder(w).Encode(customers)
	// xml.NewEncoder(w).Encode(customers)

	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}