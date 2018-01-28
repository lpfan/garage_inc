package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/d2r2/go-dht"
	"github.com/gorilla/mux"
	"github.com/stianeikeland/go-rpio"
)

var (
	pinMap = map[string]interface{}{
		"pin1": rpio.Pin(1),
		"pin2": rpio.Pin(2),
	}
)

type reqPayload struct {
	lampNumber int
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/light/turnlamp", getLamp).Methods("GET")
	router.HandleFunc("/api/v1/light/turnlamp", turnLamp).Methods("PUT")
	router.HandleFunc("/api/v1/climate", climateHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func climateHandler(w http.ResponseWriter, r *http.Request) {
	temperature, humidity, retried, err := dht.ReadDHTxxWithRetry(dht.DHT11, 4, true, 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
		temperature, humidity, retried)
}

func getLamp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test GET Method"))
}

func turnLamp(w http.ResponseWriter, r *http.Request) {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()
	decoder := json.NewDecoder(r.Body)
	var rPayload reqPayload
	err := decoder.Decode(&rPayload)

	if err != nil {
		fmt.Println(err)
	}

	defer r.Body.Close()
}
