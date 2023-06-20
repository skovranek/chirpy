package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func (cfg *apiConfig) polkaWebhooksHandler(w http.ResponseWriter, r *http.Request) {
	apiKeyAndPrefix := r.Header.Get("Authorization")
	apiKey := strings.TrimPrefix(apiKeyAndPrefix, "ApiKey ")
	if apiKey != cfg.polkaKey {
		log.Printf("Error polka API key not matching")
		respondWithError(w, http.StatusUnauthorized, "Not authorized")
		return
	}

	decoder := json.NewDecoder(r.Body)
	eventParams := struct {
		Event string `json:"event"`
		Data  struct {
			UserID int `json:"user_id"`
		} `json:"data"`
	}{}

	err := decoder.Decode(&eventParams)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode request body")
		return
	}

	if eventParams.Event != "user.upgraded" {
		emptyResp := struct{}{}
		respondWithJSON(w, http.StatusOK, emptyResp)
		return
	}

	userID := eventParams.Data.UserID
	exists, err := cfg.db.UpgradeUser(userID)
	if err != nil {
		log.Printf("Error upgrading user: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't upgrade user")
		return
	}
	if !exists {
		log.Printf("Error user does not exists, cannot upgrade user with ID #%v", userID)
		respondWithError(w, http.StatusForbidden, "Couldn't upgrade user")
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{}) // empty struct needed in response to pass tutorial test case
	return
}
