package main

import (
	"log"
	"net/http"
)

func (cfg *apiConfig) revokeRefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
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

	err = cfg.db.RevokeToken(tokenData.Token)
	if err != nil {
		log.Printf("Error revoking token: %v", err)
		respondWithError(w, http.StatusUnauthorized, "Coundn't revoke token")
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{}) // empty struct needed in response to pass tutorial test case
	return
}
