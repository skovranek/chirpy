package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (cfg *apiConfig) getChirpByIDHandler(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.GetChirps()
	if err != nil {
		log.Printf("Error getting chirps: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps")
		return
	}

	strID := chi.URLParam(r, "chirpID")
	ID, err := strconv.Atoi(strID)
	if err != nil {
		log.Printf("Error converting parameters str to int: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't convert parameters str to int")
		return
	}

	if ID > len(chirps) || ID < 1 {
		log.Printf("Error chirp ID out of range: %v", err)
		respondWithError(w, http.StatusNotFound, "Couldn't find chirp ID")
		return
	}
	chirp := chirps[ID-1]
	respondWithJSON(w, http.StatusOK, chirp)
	return
}