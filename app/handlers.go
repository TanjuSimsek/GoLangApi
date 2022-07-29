package app

import (
	"GoLangApi/service"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Customer struct {
	Name       string `json:"full_name" xml:"name"`
	City       string `json:"city" xml:"city"`
	ZipCode    string `json:"zip_code" xml:"zipcode"`
	CustomerId string `json:"customerId" xml:"customerId"`
}
type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, _ := ch.service.GetAllCustomer()
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		// json.NewEncoder(w).Encode(customers)
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
