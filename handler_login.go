package main

import (
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiConfig) loginHandler(w http.ResponseWriter, r *http.Request) {
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

	user, err := cfg.db.GetUserWithEmail(params.Email)
	if err != nil {
		log.Printf("Error getting user with email: %v", err)
		respondWithError(w, http.StatusUnauthorized, "Coundn't find email+password")
		return
	}

	pw := []byte(params.Password)
	err = bcrypt.CompareHashAndPassword(user.Password, pw)
	if err != nil {
		log.Printf("Error comparing password: %w", err)
		respondWithError(w, http.StatusUnauthorized, "Coundn't find email+password")
		return
	}

	accessToken, err := cfg.createJWT("chirpy-access", time.Hour, user.ID)
	if err != nil {
		log.Printf("Error creating access token: %v", err)
		respondWithError(w, http.StatusUnauthorized, "Coundn't create access token")
		return
	}

	sixtyDays := time.Hour * 24 * 60
	refreshToken, err := cfg.createJWT("chirpy-refresh", sixtyDays, user.ID)
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
		Token:  accessToken,
		RefreshToken: refreshToken,
	}

	respondWithJSON(w, http.StatusOK, respBody)
	return
}
