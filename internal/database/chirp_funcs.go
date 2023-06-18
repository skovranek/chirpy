package database

import (
	"log"
)

func (db *DB) CreateChirp(authorID int, body string) (Chirp, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return Chirp{}, err
	}

	chirpID := len(dbStruct.Chirps)
	for {
		chirpID++
		if _, exists := dbStruct.Chirps[chirpID]; exists {
			log.Printf("Chirp ID already assigned: #%v\n", chirpID)
		} else {
			break
		}
	}
	chirp := Chirp{
		ID:       chirpID,
		AuthorID: authorID,
		Body:     body,
	}

	dbStruct.Chirps[chirpID] = chirp

	err = db.writeDB(dbStruct)
	if err != nil {
		return Chirp{}, err
	}

	return chirp, nil
}

func (db *DB) DeleteChirp(chirpID int) error {
	dbStruct, err := db.loadDB()
	if err != nil {
		return err
	}

	delete(dbStruct.Chirps, chirpID)

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