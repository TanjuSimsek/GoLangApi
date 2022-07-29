package app

import (
	"GoLangApi/domain"
	"GoLangApi/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	//Wiring

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	//
	router.HandleFunc("/getAllCustomers", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
