package main

import (
	"log"
	"net/http"
    "service_b_rest/api/delayed"
	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/delayed", delayed.DelayedResponse)
	log.Fatal(http.ListenAndServe(":8080", router))
}
