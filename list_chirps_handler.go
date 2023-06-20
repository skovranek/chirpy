package main

import (
	"log"
	"net/http"
	"sort"
	"strconv"
)

type Chirp struct {
	ID       int    `json:"id"`
	AuthorID int    `json:"author_id"`
	Body     string `json:"body"`
}

func (cfg *apiConfig) listChirpsHandler(w http.ResponseWriter, r *http.Request) {
	unsortedChirps, err := cfg.db.GetChirps()
	if err != nil {
		log.Printf("Error getting chirps: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps")
		return
	}

	authorID := 0
	authorIDStr := r.URL.Query().Get("author_id")
	if authorIDStr != "" {
		authorID, err = strconv.Atoi(authorIDStr)
		if err != nil {
			log.Printf("Error converting ID str to int: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Couldn't get author ID")
			return
		}
	}

	respChirps := []Chirp{}
	for _, chirp := range unsortedChirps {
		if authorIDStr == "" || chirp.AuthorID == authorID {
			respChirps = append(respChirps, Chirp{
				ID:       chirp.ID,
				AuthorID: chirp.AuthorID,
				Body:     chirp.Body,
			})
		}
	}

	sortParam := r.URL.Query().Get("sort")
	if sortParam == "desc" {
		sort.Slice(respChirps, func(i, j int) bool {
			return respChirps[i].ID > respChirps[j].ID
		})
	} else {
		sort.Slice(respChirps, func(i, j int) bool {
			return respChirps[i].ID < respChirps[j].ID
		})
	}

	respondWithJSON(w, http.StatusOK, respChirps)
	return
}
