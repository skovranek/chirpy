package main

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiConfig) loginHandler(w http.ResponseWriter, r *http.Request) {
	userParams, err := getUserParams(r)
	if err != nil {
		log.Printf("Error decoding user parameters: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't get request body")
		return
	}

	user, err := cfg.db.GetUserByEmail(userParams.Email)
	if err != nil {
		log.Printf("Error getting user with email: %v", err)
		respondWithError(w, http.StatusUnauthorized, "Coundn't find email+password")
		return
	}

	pw := []byte(userParams.Password)
	err = bcrypt.CompareHashAndPassword(user.Password, pw)
	if err != nil {
		log.Printf("Error comparing password: %v", err)
		respondWithError(w, http.StatusUnauthorized, "Coundn't find email+password")
		return
	}

	accessToken, err := cfg.createJWT("chirpy-access", cfg.accessJWTExpInHours, user.ID)
	if err != nil {
		log.Printf("Error creating access token: %v", err)
		respondWithError(w, http.StatusUnauthorized, "Coundn't create access token")
		return
	}

	refreshToken, err := cfg.createJWT("chirpy-refresh", cfg.refreshJWTExpInHours, user.ID)
	if err != nil {
		log.Printf("Error creating refresh token: %v", err)
		respondWithError(w, http.StatusUnauthorized, "Coundn't create refresh token")
		return
	}

	respBody := struct {
		ID           int    `json:"id"`
		Email        string `json:"email"`
		IsChirpyRed  bool   `json:"is_chirpy_red"`
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}{
		ID:           user.ID,
		Email:        user.Email,
		IsChirpyRed:  user.IsChirpyRed,
		Token:        accessToken,
		RefreshToken: refreshToken,
	}

	respondWithJSON(w, http.StatusOK, respBody)
	return
}
