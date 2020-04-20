package main

import (
	"log"
	"net/http"
	"service_b_rest/api/delayed"
	"service_b_rest/api/health"

	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/delayed", delayed.DelayedResponse)
	router.HandleFunc("/health",health.Health)
	log.Fatal(http.ListenAndServe(":8080", router))
}
