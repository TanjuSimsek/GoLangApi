package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name       string `json:"full_name" xml:"name"`
	City       string `json:"city" xml:"city"`
	ZipCode    string `json:"zip_code" xml:"zipcode"`
	CustomerId string `json:"customerId" xml:"customerId"`
}

func getApi(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Apiyi Tetiklediniz...")

}
func getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers := []Customer{

		{Name: "TANJU", City: "iSTANBUL", ZipCode: "34000"},
		{Name: "Elif", City: "iSTANBUL", ZipCode: "34000"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		// json.NewEncoder(w).Encode(customers)
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
func getCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	// fmt.Println(vars["customer_id"])
	// fmt.Fprint(w, vars["customer_id"])
	customers := []Customer{

		{Name: "TANJU", City: "iSTANBUL", ZipCode: "34000", CustomerId: "1"},
		{Name: "Elif", City: "iSTANBUL", ZipCode: "34000", CustomerId: "2"},
	}
	var customer Customer
	for _, v := range customers {
		if v.CustomerId == vars["customer_id"] {
			// Found!
			customer.City = v.City
			customer.CustomerId = v.CustomerId
			customer.ZipCode = v.ZipCode
			customer.Name = v.Name

		}
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		// json.NewEncoder(w).Encode(customers)
		if (customer == Customer{}) {
			var message = "Bu id ye ait müşteri yok" + vars["customer_id"]
			xml.NewEncoder(w).Encode(message)

		} else {
			xml.NewEncoder(w).Encode(customer)
		}

	} else {
		if (customer == Customer{}) {
			var message = "Bu id ye ait müşteri yok" + vars["customer_id"]
			json.NewEncoder(w).Encode(message)
		} else {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(customer)
		}

	}
	// customers := []Customer{

	// 	{Name: "TANJU", City: "iSTANBUL", ZipCode: "34000"},
	// 	{Name: "Elif", City: "iSTANBUL", ZipCode: "34000"},
	// }

	// if r.Header.Get("Content-Type") == "application/xml" {
	// 	w.Header().Add("Content-Type", "application/xml")
	// 	// json.NewEncoder(w).Encode(customers)
	// 	xml.NewEncoder(w).Encode(customers)
	// } else {
	// 	w.Header().Add("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(customers)
	// }

}
func createCustomer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "post method trigered")
}
func getAllCustomersXml(w http.ResponseWriter, r *http.Request) {

	customers := []Customer{

		{Name: "TANJU", City: "iSTANBUL", ZipCode: "34000"},
		{Name: "Elif", City: "iSTANBUL", ZipCode: "34000"},
	}
	w.Header().Add("Content-Type", "application/xml")
	// json.NewEncoder(w).Encode(customers)
	xml.NewEncoder(w).Encode(customers)

}
