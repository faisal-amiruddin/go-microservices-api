package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/faisal-amiruddin/go-microservices-api/domain"
	"github.com/faisal-amiruddin/go-microservices-api/service"
	"github.com/gorilla/mux"
)

func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)

	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	//start server
	fmt.Println("Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}