package database

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

func NewDB(path string) (*DB, error) {
	db := DB{
		path: path,
		mux:  &sync.RWMutex{},
	}
	err := db.ensureDB()
	if err != nil {
		return &DB{}, err
	}

	return &db, nil
}

func (db *DB) ensureDB() error {
	db.mux.Lock()
	defer db.mux.Unlock()

	f, err := os.OpenFile(db.path, os.O_RDWR|os.O_CREATE, 0600) // 0600 is user read & write
	defer f.Close()
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) loadDB() (DBStruct, error) {
	db.mux.RLock()
	defer db.mux.RUnlock()

	f, err := os.OpenFile(db.path, os.O_RDONLY, 0400) // 0400 is user read only
	defer f.Close()
	if err != nil {
		return DBStruct{}, err
	}

	fileInfo, err := f.Stat()
	if err != nil {
		return DBStruct{}, err
	}

	if fileInfo.Size() == 0 {
		dbStruct := DBStruct{
			Users:         map[int]User{},
			Emails:        map[string]int{},
			RevokedTokens: map[string]time.Time{},
			Chirps:        map[int]Chirp{},
		}
		return dbStruct, nil
	}

	dat := make([]byte, fileInfo.Size())
	f.Read(dat)
	if err != nil {
		return DBStruct{}, err
	}

	dbStruct := DBStruct{}
	err = json.Unmarshal(dat, &dbStruct)
	if err != nil {
		return DBStruct{}, err
	}
	return dbStruct, nil
}

func (db *DB) writeDB(dbStruct DBStruct) error {
	db.mux.Lock()
	defer db.mux.Unlock()

	dat, err := json.Marshal(&dbStruct)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(db.path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600) // 0200 is user write only
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.Write(dat)
	if err != nil {
		return err
	}
	return nil
}

func RemoveFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
