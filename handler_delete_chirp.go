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
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate token")
		return
	}

	chirps, err := cfg.db.GetChirps()
	if err != nil {
		log.Printf("Error getting chirps: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps")
		return
	}

	chirpIDStr := chi.URLParam(r, "chirpID")
	chirpID, err := strconv.Atoi(chirpIDStr)
	if err != nil {
		log.Printf("Error converting parameters str to int: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't convert parameters str to int")
		return
	}

	if chirpID > len(chirps) || chirpID < 1 {
		log.Printf("Error chirp ID out of range: %v", err)
		respondWithError(w, http.StatusNotFound, "Couldn't find chirp ID")
		return
	}
	chirp := chirps[chirpID-1]
	if chirp.AuthorID != tokenData.userID {
		log.Print("Error user not author and is not authorized to delete chirp")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = cfg.db.DeleteChirp(chirpID)
	if err != nil {
		log.Printf("Error deleting chirp from database: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Coundn't delete chirp from database")
		return
	}

	emptyResp := struct{}{}
	respondWithJSON(w, http.StatusOK, emptyResp)
	return
}
