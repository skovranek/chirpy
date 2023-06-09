package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenData struct {
	Token  string
	userID int
	issuer string
}

func (cfg *apiConfig) getJWTData(r *http.Request) (TokenData, error) {
	tokenAndPrefixStr := r.Header.Get("Authorization")
	unparsedTokenStr := strings.TrimPrefix(tokenAndPrefixStr, "Bearer ")

	token, err := jwt.ParseWithClaims(unparsedTokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return cfg.jwtSecret, nil
	})
	if err != nil {
		err = fmt.Errorf("Error parsing token: %w", err)
		return TokenData{}, err
	}

	claims, exists := token.Claims.(*jwt.RegisteredClaims)
	if !exists {
		err = fmt.Errorf("Error validating token: token does not exist")
		return TokenData{}, err
	}
	if !token.Valid {
		err = fmt.Errorf("Error token is invalid")
		return TokenData{}, err
	}

	expiresAt, err := claims.GetExpirationTime()
	if expiresAt.Time.Before(time.Now()) || err != nil {
		err = fmt.Errorf("Error token is expired: %w", err)
		return TokenData{}, err
	}

	userIDStr, err := claims.GetSubject()
	if err != nil {
		err = fmt.Errorf("Error validating ID: %w", err)
		return TokenData{}, err
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		err = fmt.Errorf("Error converting id str to int: %w", err)
		return TokenData{}, err
	}

	issuer, err := claims.GetIssuer()
	if err != nil {
		err = fmt.Errorf("Error validating ID: %w", err)
		return TokenData{}, err
	}

	tokenData := TokenData{
		Token:  unparsedTokenStr,
		userID: userID,
		issuer: issuer,
	}
	return tokenData, nil
}
