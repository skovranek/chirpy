package main

import (
	"log"
	"net/http"
	"time"
)

func (cfg *apiConfig) refreshAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	tokenData, err := cfg.getJWTData(r)
	if err != nil {
		log.Printf("Error getting token data: %v", err)
		respondWithError(w, http.StatusUnauthorized, "Coundn't get token data")
		return
	}
	if tokenData.issuer != "chirpy-refresh" {
		log.Printf("Error validating token: must be a refresh token")
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate token: must be a refresh token")
		return
	}

	accessToken, err := cfg.createJWT("chirpy-access", time.Hour, tokenData.userID)
	if err != nil {
		log.Printf("Error refreshing token: %v", err)
		respondWithError(w, http.StatusUnauthorized, "Coundn't refresh token")
		return
	}

	respBody := struct {
		Token string
	}{
		Token: accessToken,
	}

	respondWithJSON(w, http.StatusOK, respBody)
	return
}
