package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// The dispatcher function will read from this channel
var dataChan chan map[string]string

const webhookURL = "https://webhook.site/your-webhook-endpoint"

func main() {
	// set up dispatcher func and the channel
	dataChan = make(chan map[string]string, 100)
	go dispatch()

	// Server and router stuff
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", requestHandler)
	http.ListenAndServe(":8080", mux)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var incomingPayload map[string]string
	json.NewDecoder(r.Body).Decode(&incomingPayload)
	dataChan <- incomingPayload
	w.WriteHeader(http.StatusAccepted)
	fmt.Println(incomingPayload)
}
