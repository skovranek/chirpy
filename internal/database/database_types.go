package database

import (
	"sync"
	"time"
)

type DB struct {
	path string
	mux  *sync.RWMutex
}

type DBStruct struct {
	Chirps        map[int]Chirp        `json:"chirps"`
	Users         map[int]User         `json:"users"`
	RevokedTokens map[string]time.Time `json:"revoked_tokens"`
}

type Chirp struct {
	ID       int    `json:"id"`
	AuthorID int    `json:"author_id"`
	Body     string `json:"body"`
}

type User struct {
	ID           int    `json:"id,omitempty"`
	Password     []byte `json:"password,omitempty"`
	Email        string `json:"email,omitempty"`
	IsChirpyRed  bool   `json:"is_chirpy_red"`
}