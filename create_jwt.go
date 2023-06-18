package main

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (cfg *apiConfig) createJWT(issuer string, duration time.Duration, userID int) (string, error) {
	claims := &jwt.RegisteredClaims{
		Issuer:    issuer,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		Subject:   strconv.Itoa(userID),
	}

	unsignedTokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedTokenString, err := unsignedTokenString.SignedString(cfg.jwtSecret)
	if err != nil {
		return "", err
	}

	return signedTokenString, nil
}
