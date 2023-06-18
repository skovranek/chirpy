package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/skovranek/chirpy/internal/database"
)

func configure() apiConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(`Error loading ".env" file`)
	}
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	polkaKey := os.Getenv("POLKA_API_KEY")

	dbg := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	if *dbg {
		err := database.RemoveFile(DB_PATH)
		if err != nil {
			log.Printf(`Error removing file "%s": %v`, DB_PATH, err)
		}
	}

	db, err := database.NewDB(DB_PATH)
	if err != nil {
		log.Printf("Error creating type DB at '%s': %v", DB_PATH, err)
	}

	return apiConfig{
		db:        db,
		jwtSecret: jwtSecret,
		polkaKey:  polkaKey,
	}
}

type apiConfig struct {
	db             *database.DB
	jwtSecret      []byte
	polkaKey       string
	fileServerHits int
}