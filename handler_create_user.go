package main

import (
	"log"
	"net/http"
)

func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
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

	user, err := cfg.db.CreateUser(params.Password, params.Email)
	if err != nil {
		log.Printf("Error creating user in database: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Coundn't create user in database")
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
	return
}