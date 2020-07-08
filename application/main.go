package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/temperatures", getAllTemperatures).Methods("GET")
	router.HandleFunc("/temperature", postTemperature).Methods("POST")
	log.Fatal(http.ListenAndServe(":4444", router))
}
