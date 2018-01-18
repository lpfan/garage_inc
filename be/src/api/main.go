package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/light/turnlamp", GetLamp).Methods("GET")
	router.HandleFunc("/api/v1/light/turnlamp", TurnLamp).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetLamp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test GET Method"))
}

func TurnLamp(w http.ResponseWriter, r *http.Request) {}
