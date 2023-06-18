package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, statusCode int, errMsg string) {
	respBody := struct {
		Error string `json:"error"`
	}{
		Error: errMsg,
	}
	data, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
	return
}
