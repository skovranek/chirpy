package database

import (
	"fmt"
	"time"
)

func (db *DB) GetUserWithEmail(email string) (User, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	user, exists := dbStruct.getUserByEmail(email)
	if !exists {
		err := fmt.Errorf("Error email+password not found in users: %s", email)
		return User{}, err
	}

	return user, nil
}

func (db *DB) Revoke(token string) error {
	dbStruct, err := db.loadDB()
	if err != nil {
		return err
	}

	_, isRevoked := dbStruct.RevokedTokens[token]
	if isRevoked {
		return fmt.Errorf("Error refresh token is already revoked")
	}

	dbStruct.RevokedTokens[token] = time.Now()

	err = db.writeDB(dbStruct)
	if err != nil {
		return err
	}

	return nil
}
