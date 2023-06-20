package main

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiConfig) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	userParams, err := getUserParams(r)
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

	newPassword := []byte(userParams.Password)
	hashedPW, err := bcrypt.GenerateFromPassword(newPassword, bcrypt.MinCost) // cost param min-max: 4-31
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Coundn't update user in database")
		return
	}

	user, err := cfg.db.UpdateUser(tokenData.userID, hashedPW, userParams.Email)
	if err != nil {
		log.Printf("Error updating user in database: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Coundn't update user in database")
		return
	}

	respondWithJSON(w, http.StatusOK, user)
	return
}
