package database

import (
	"fmt"
	"log"
	"time"
)

func (db *DB) CreateUser(pw []byte, email string) (User, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	_, emailUsed := dbStruct.Emails[email]
	if emailUsed {
		err := fmt.Errorf("Error email already used: %s", email)
		return User{}, err
	}

	uniqueID := len(dbStruct.Users)
	for {
		uniqueID++
		if _, exists := dbStruct.Users[uniqueID]; exists {
			log.Printf("User ID already assigned: #%v\n", uniqueID)
		} else {
			break
		}
	}

	user := User{
		ID:       uniqueID,
		Password: pw,
		Email:    email,
	}

	dbStruct.Users[user.ID] = user
	dbStruct.Emails[email] = user.ID

	err = db.writeDB(dbStruct)
	if err != nil {
		return User{}, err
	}

	respUser := User{
		ID:    uniqueID,
		Email: email,
	}
	return respUser, nil
}

func (db *DB) UpdateUser(id int, pw []byte, email string) (User, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	IDofUserWithEmail, emailUsed := dbStruct.Emails[email]
	if id != IDofUserWithEmail && emailUsed {
		err := fmt.Errorf("Error email already used: %s", email)
		return User{}, err
	}

	user, userExists := dbStruct.Users[id]
	if !userExists {
		err := fmt.Errorf("Error email+password not found in users: %s", email)
		return User{}, err
	}

	if email != user.Email {
		delete(dbStruct.Emails, user.Email)
		dbStruct.Emails[email] = id
		user.Email = email
	}

	user.Password = pw
	dbStruct.Users[id] = user

	err = db.writeDB(dbStruct)
	if err != nil {
		return User{}, err
	}

	respUser := User{
		ID:    id,
		Email: email,
	}
	return respUser, nil
}

func (db *DB) GetUserByEmail(email string) (User, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	userID, emailUsed := dbStruct.Emails[email]
	if !emailUsed {
		err := fmt.Errorf("Error email+password not found in users: %s", email)
		return User{}, err
	}
	user, userExists := dbStruct.Users[userID]
	if !userExists {
		err := fmt.Errorf("Error email+password not found in users: %s", email)
		return User{}, err
	}

	return user, nil
}

func (db *DB) RevokeToken(token string) error {
	dbStruct, err := db.loadDB()
	if err != nil {
		return err
	}

	if _, isRevoked := dbStruct.RevokedTokens[token]; isRevoked {
		return fmt.Errorf("Error token is already revoked")
	}

	dbStruct.RevokedTokens[token] = time.Now()

	err = db.writeDB(dbStruct)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UpgradeUser(id int) (bool, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return false, err
	}

	user, exists := dbStruct.Users[id]
	if !exists {
		return false, nil
	}

	user.IsChirpyRed = true
	dbStruct.Users[id] = user

	err = db.writeDB(dbStruct)
	if err != nil {
		return false, err
	}

	return true, nil
}
