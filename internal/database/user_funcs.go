package database

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (db *DB) CreateUser(pw string, email string) (User, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	_, exists := dbStruct.getUserByEmail(email)
	if exists {
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

	hashedPW, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost) // cost param min/max: 4-31
	if err != nil {
		return User{}, err
	}

	user := User{
		ID:       uniqueID,
		Password: hashedPW,
		Email:    email,
	}

	dbStruct.Users[user.ID] = user

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

func (db *DB) UpdateUser(id int, pw, email string) (User, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	hashedPW, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost) // cost param min/max: 4-31
	if err != nil {
		return User{}, err
	}

	userWithEmail, exists := dbStruct.getUserByEmail(email)
	if userWithEmail.ID != id && exists {
		err := fmt.Errorf("Error email already used: %s", email)
		return User{}, err
	}

	user := dbStruct.Users[id]
	user.Password = hashedPW
	user.Email = email
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
		return true, err
	}

	return true, nil
}