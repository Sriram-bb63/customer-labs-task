package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const webhookURL = "https://webhook.site/"

// The dispatcher goroutine will read from this channel
var dataChan chan map[string]string

func main() {
	log.Println("Starting server")

	dataChan = make(chan map[string]string, 10000)
	log.Println("Channel created")

	for range 10000 {
		go dispatch()
	}
	log.Println("Goroutine started")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", getHandler)
	mux.HandleFunc("POST /", postHandler)
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("GET request received from ", r.RemoteAddr)
	w.Write([]byte("Server status - running"))
	log.Println("GET request handled successfully")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("POST request received from ", r.RemoteAddr)
	defer r.Body.Close()
	var incomingPayload map[string]string
	err := json.NewDecoder(r.Body).Decode(&incomingPayload)
	if err != nil {
		log.Println("Error decoding JSON payload: ", err)
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	if len(dataChan) == cap(dataChan) {
		log.Println("dataChan is full, dropping the current request")
		http.Error(w, "Server is busy", http.StatusServiceUnavailable)
		return
	}
	dataChan <- incomingPayload
	log.Println("Payload sent to dataChan")
	log.Println("POST request handled successfully")
}
