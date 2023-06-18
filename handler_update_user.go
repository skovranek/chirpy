package main

import (
	"log"
	"net/http"
)

func (cfg *apiConfig) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	params := struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}{}

	err := getUserParams(r, &params)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode request body")
		return
	}

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

	user, err := cfg.db.UpdateUser(tokenData.userID, params.Password, params.Email)
	if err != nil {
		log.Printf("Error updating user in database: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Coundn't update user in database")
		return
	}

	respondWithJSON(w, http.StatusOK, user)
	return
}
