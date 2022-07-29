package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// mux := http.NewServeMux() ->>>Mux

	router := mux.NewRouter()

	// http.HandleFunc("/tanju", getApi)
	// http.HandleFunc("/getAllCustomers", getAllCustomers)

	router.HandleFunc("/tanju", getApi).Methods(http.MethodGet)
	router.HandleFunc("/getAllCustomers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/getAllCustomers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/getAllCustomers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// http.HandleFunc("/getAllCustomersXml", getAllCustomersXml)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
