package app

import (
	"capi/domain"
	"capi/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// * create ServeRoute
	mux := mux.NewRouter()

	// * defining routes// * wiring
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	mux.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)


	// * starting the server
	fmt.Println("starting the server localhost:9000")
	http.ListenAndServe(":9000", mux)
}