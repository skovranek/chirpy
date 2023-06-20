package database

import (
	"fmt"
	"log"
)

func (db *DB) CreateChirp(authorID int, body string) (Chirp, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return Chirp{}, err
	}

	uniqueID := len(dbStruct.Chirps)
	for {
		uniqueID++
		if _, exists := dbStruct.Chirps[uniqueID]; exists {
			log.Printf("Chirp ID already assigned: #%v\n", uniqueID)
		} else {
			break
		}
	}
	chirp := Chirp{
		ID:       uniqueID,
		AuthorID: authorID,
		Body:     body,
	}

	dbStruct.Chirps[uniqueID] = chirp

	err = db.writeDB(dbStruct)
	if err != nil {
		return Chirp{}, err
	}

	return chirp, nil
}

func (db *DB) DeleteChirp(userID, chirpID int) error {
	dbStruct, err := db.loadDB()
	if err != nil {
		return err
	}

	if chirp, exists := dbStruct.Chirps[chirpID]; !exists {
		err := fmt.Errorf("Error cannot get chirp with ID: #%v", chirpID)
		return err
	} else if chirp.AuthorID == userID {
		delete(dbStruct.Chirps, chirpID)
	} else {
		err := fmt.Errorf("Error user #%v did not author chirp #%v", userID, chirpID)
		return err
	}

	err = db.writeDB(dbStruct)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetChirps() ([]Chirp, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return []Chirp{}, err
	}

	chirps := []Chirp{}
	for _, chirp := range dbStruct.Chirps {
		chirps = append(chirps, chirp)
	}

	return chirps, nil
}

func (db *DB) GetChirpByID(id int) (Chirp, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return Chirp{}, err
	}

	chirp, exists := dbStruct.Chirps[id]
	if !exists {
		err = fmt.Errorf("Error cannot get chirp with ID: #%v", id)
		return Chirp{}, err
	}

	return chirp, nil
}
