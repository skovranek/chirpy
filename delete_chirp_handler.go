package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (cfg *apiConfig) deleteChirpHandler(w http.ResponseWriter, r *http.Request) {
	tokenData, err := cfg.getJWTData(r)
	if err != nil {
		log.Printf("Error getting token data: %v", err)
		respondWithError(w, http.StatusUnauthorized, "Coundn't get token data")
		return
	}
	if tokenData.issuer != "chirpy-access" {
		log.Printf("Error validating token: must be an access token")
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate token: must be an access token")
		return
	}

	chirpIDStr := chi.URLParam(r, "chirpID")
	chirpID, err := strconv.Atoi(chirpIDStr)
	if err != nil {
		log.Printf("Error converting ID str to int: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirp ID")
		return
	}

	err = cfg.db.DeleteChirp(tokenData.userID, chirpID)
	if err != nil {
		log.Printf("Error deleting chirp from database: %v", err)
		respondWithError(w, http.StatusForbidden, "Coundn't delete chirp from database")
		return
	}

	emptyResp := struct{}{}
	respondWithJSON(w, http.StatusOK, emptyResp)
	return
}
