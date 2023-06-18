package main

import (
	"log"
	"net/http"
)

func (cfg *apiConfig) revokeHandler(w http.ResponseWriter, r *http.Request) {
	tokenData, err := cfg.getJWTData(r)
	if err != nil {
		log.Printf("Error getting token data: %v", err)
		respondWithError(w, http.StatusUnauthorized, "Coundn't get token data")
		return
	}
	if tokenData.issuer != "chirpy-refresh" {
		log.Printf("Error validating token: must be a refresh token")
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate token")
		return
	}

	err = cfg.db.Revoke(tokenData.unparsedTokenStr)
	if err != nil {
		log.Printf("Error revoking token: %v", err)
		respondWithError(w, http.StatusUnauthorized, "Coundn't revoke token")
		return
	}

	emptyResp := struct{}{}
	respondWithJSON(w, http.StatusOK, emptyResp)
	return
}
