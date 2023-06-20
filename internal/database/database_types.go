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
	Users         map[int]User         `json:"users"`
	Emails        map[string]int       `json:"emails"`
	RevokedTokens map[string]time.Time `json:"revoked_tokens"`
	Chirps        map[int]Chirp        `json:"chirps"`
}

type User struct {
	ID          int    `json:"id,omitempty"`
	Password    []byte `json:"password,omitempty"`
	Email       string `json:"email,omitempty"`
	IsChirpyRed bool   `json:"is_chirpy_red"`
}

type Chirp struct {
	ID       int    `json:"id"`
	AuthorID int    `json:"author_id"`
	Body     string `json:"body"`
}
