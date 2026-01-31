package app

import (
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()

	//define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomer)

	//start server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}