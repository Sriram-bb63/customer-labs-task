package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", requestHandler)
	http.ListenAndServe(":8080", mux)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var incomingPayload map[string]string
	json.NewDecoder(r.Body).Decode(&incomingPayload)
	dispatch(incomingPayload)
	w.WriteHeader(http.StatusAccepted)
	fmt.Println(incomingPayload)
}
