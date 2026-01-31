package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name	string	`json:"full_name"`
	City	string	`json:"city"`
	Zipcode	string	`json:"zip_code"`
}

func main() {
	//define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomer)

	//start server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers := []Customer {
		{ "Naruto", "Konoha", "12345" },
		{ "Sasuke", "konoha", "12345" },
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(customers)
}