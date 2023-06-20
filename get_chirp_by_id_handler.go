package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (cfg *apiConfig) getChirpByIDHandler(w http.ResponseWriter, r *http.Request) {
	chirpIDStr := chi.URLParam(r, "chirpID")
	chirpID, err := strconv.Atoi(chirpIDStr)
	if err != nil {
		log.Printf("Error converting ID str to int: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirp ID")
		return
	}

	chirp, err := cfg.db.GetChirpByID(chirpID)
	if err != nil {
		log.Printf("Error getting chirp: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirp")
		return
	}

	respondWithJSON(w, http.StatusOK, chirp)
	return
}
