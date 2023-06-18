package main

import (
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/skovranek/chirpy/internal/database"
)

func (cfg *apiConfig) getChirpsHandler(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.GetChirps()
	if err != nil {
		log.Printf("Error getting chirps: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps")
		return
	}

	authorIDStr := r.URL.Query().Get("author_id")
	if len(authorIDStr) > 0 {
		authorID, err := strconv.Atoi(authorIDStr)
		if err != nil {
			log.Printf("Error converting id str to int: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Couldn't get author ID")
			return
		}

		chirpsByAuthor := []database.Chirp{}
		for _, chirp := range chirps {
			if chirp.AuthorID == authorID {
				chirpsByAuthor = append(chirpsByAuthor, chirp)
			}
		}
		chirps = chirpsByAuthor
	}

	sortParam := r.URL.Query().Get("sort")
	if sortParam == "desc" {
		sort.Slice(chirps, func(i, j int) bool { return chirps[i].ID > chirps[j].ID })
	} else {
		sort.Slice(chirps, func(i, j int) bool { return chirps[i].ID < chirps[j].ID })
	}

	respondWithJSON(w, http.StatusOK, chirps)
	return
}