package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func (cfg *apiConfig) createChirpHandler(w http.ResponseWriter, r *http.Request) {
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

	decoder := json.NewDecoder(r.Body)
	chirpParams := struct {
		Body string `json:"body"`
	}{}
	err = decoder.Decode(&chirpParams)
	if err != nil {
		log.Printf("Error decoding parameters: %v", err)
		respondWithError(w, http.StatusUnauthorized, "Couldn't get request body")
		return
	}

	if len(chirpParams.Body) > 140 {
		log.Printf("Chirp is too long")
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}
	cleanedBody := censorProfanity(chirpParams.Body)
	chirp, err := cfg.db.CreateChirp(tokenData.userID, cleanedBody)
	if err != nil {
		log.Printf("Error saving chirp to database: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Coundn't save chirp to database")
		return
	}

	respondWithJSON(w, http.StatusCreated, chirp)
	return
}

func censorProfanity(str string) string {
	fakeProfanity := ".kerfuffle.sharbert.fornax."
	words := strings.Split(str, " ")
	if len(words) > 0 {
		for i, word := range words {
			if strings.Contains(fakeProfanity, "."+strings.ToLower(word)+".") {
				words[i] = "****"
			}
		}
	}
	return strings.Join(words, " ")
}
