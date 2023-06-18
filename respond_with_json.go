package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
	return
}
