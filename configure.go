package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"

	"github.com/skovranek/chirpy/internal/database"
)

/*
EXAMPLE '.env' file:
ROOT=.
PORT=8080
DB_PATH=database.json
POLKA_API_KEY=f271c81ff7084ee5b99a5091b42d486e
JWT_SECRET=CKpJVLHqOoYtKX/hjkQ6iPtVhqeqmAKYF4uPfqGoQxTVVe8ZMbedqRcjUrhlkiy1keNbSQq3Cn9RnZ2xTKM8GA==
ACCESS_JWT_EXP_IN_HOURS=1
REFRESH_JWT_EXP_IN_HOURS=1440
*/

type apiConfig struct {
	root                 string
	port                 string
	db_path              string
	db                   *database.DB
	polkaKey             string
	jwtSecret            []byte
	accessJWTExpInHours  time.Duration
	refreshJWTExpInHours time.Duration
	fileServerHits       int
}

func configure() *apiConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(`Error loading ".env" file`)
	}
	root := os.Getenv("ROOT")
	port := os.Getenv("PORT")
	db_path := os.Getenv("DB_PATH")
	polkaKey := os.Getenv("POLKA_API_KEY")
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	accessJWTExpStr := os.Getenv("ACCESS_JWT_EXP_IN_HOURS")
	refreshJWTExpStr := os.Getenv("REFRESH_JWT_EXP_IN_HOURS")

	if len(root) == 0 {
		log.Printf("Error 'ROOT' env_var not found")
		root = "."
		log.Printf("Default used for root filepath: %s", root)
	}
	if len(port) == 0 {
		log.Printf("Error 'PORT' env_var not found")
		port = "8080"
		log.Printf("Default used for port num: %s", port)
	}
	if len(db_path) == 0 {
		log.Printf("Error 'DB_PATH' env_var not found")
		db_path = "database.json"
		log.Printf("Default used for database filepath: %s", port)
	}

	dbg := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	if *dbg {
		err := database.RemoveFile(db_path)
		if err != nil {
			log.Printf(`Error removing file "%s": %v`, db_path, err)
		}
	}

	db, err := database.NewDB(db_path)
	if err != nil {
		log.Printf("Error creating type DB at '%s': %v", db_path, err)
	}

	accessJWTExpInt, err := strconv.Atoi(accessJWTExpStr)
	if err != nil {
		log.Printf("Error converting accessJWTExpStr to int: %v", err)
		accessJWTExpInt = 1
		log.Printf("Default value used for access JWT Expiration: %v hour", accessJWTExpInt)
	}
	accessJWTExpInHours := time.Hour * time.Duration(accessJWTExpInt)

	refreshJWTExpInt, err := strconv.Atoi(refreshJWTExpStr)
	if err != nil {
		log.Printf("Error converting accessJWTExpStr to int: %v", err)
		refreshJWTExpInt = 1440
		log.Printf("Default value used for refresh JWT Expiration: %v hours", refreshJWTExpInt)
	}
	refreshJWTExpInHours := time.Hour * time.Duration(refreshJWTExpInt)

	//log.Printf("ENV_VARS:\nROOT: %s, \nPORT: %s, \nDP_PATH: %s, \nACCESS_JWT_EXP_IN_HOURS: %v, \nREFRESH_JWT_EXP_IN_HOURS: %v", root, port, db_path, accessJWTExpInt, refreshJWTExpInt)

	return &apiConfig{
		root:                 root,
		port:                 port,
		db_path:              db_path,
		db:                   db,
		jwtSecret:            jwtSecret,
		accessJWTExpInHours:  accessJWTExpInHours,
		refreshJWTExpInHours: refreshJWTExpInHours,
		polkaKey:             polkaKey,
	}
}
