package main

import "capi/app"

// type Customer struct {
// 	Name string `json:"name" xml:"name"`
// 	City string `json:"city" xml:"city"`
// 	ZipCode string `json:"zip_code" xml:"zipCode"`
// }

// var customers []Customer = []Customer{
// 	{"Bruyne", "Banten", "15540"},
// 	{"Brown", "Surabaya", "15560"},
// }

func main() {

	// * defining routes
	// http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello Celerates!")
	// })
	// http.HandleFunc("/greet", greet)
	// http.HandleFunc("/customers", getCustomers)

	// // * starting the server
	// fmt.Println("starting the server localhost:9000")
	// http.ListenAndServe(":9000", nil)

	app.Start()
}

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello Celerates!")
// }

// func getCustomers(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprint(w, "Get Custoem Endpoint!")
	
// 	// w.Header().Add("Content Type", "application/json")
// 	// json.NewEncoder(w).Encode(customers)
// 	// xml.NewEncoder(w).Encode(customers)

// 	if r.Header.Get("Content-Type") == "application/xml" {
// 		w.Header().Add("Content-Type", "application/xml")
// 		xml.NewEncoder(w).Encode(customers)
// 	}else{
// 		w.Header().Add("Content-Type", "application/xml")
// 		xml.NewEncoder(w).Encode(customers)
// 	}
// }
