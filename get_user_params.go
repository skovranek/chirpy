package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getUserParams(r *http.Request) (UserParams, error) {
	decoder := json.NewDecoder(r.Body)
	userParams := UserParams{}

	err := decoder.Decode(&userParams)
	if err != nil {
		err = fmt.Errorf("Error decoding user parameters: %w", err)
		return UserParams{}, err
	}

	return userParams, nil
}

type UserParams struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
