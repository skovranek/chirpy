package main

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	userParams, err := getUserParams(r)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't get request body")
		return
	}

	newPassword := []byte(userParams.Password)
	hashedPW, err := bcrypt.GenerateFromPassword(newPassword, bcrypt.MinCost) // cost param min-max: 4-31
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Coundn't create user in database")
		return
	}

	user, err := cfg.db.CreateUser(hashedPW, userParams.Email)
	if err != nil {
		log.Printf("Error creating user in database: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Coundn't create user in database")
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
	return
}
