package main

import (
	"log"
	"net/http"
)

/*
TODO:
[x] fix lack of auth headers causing index panic when len called on missing param string
[x] get rid of get user by ID, redundant since maps return key
[ ] refactor
[ ] extract
[x] extract server
[x] extract auth
[ ] extract in db
[ ] remove JWT and bcrypt from database
[ ] errors fmt
[ ] add created_at time field to chirps?
[ ] sort by created_at?
[ ] add UUIDs ?
[ ] add mux Rlock defer Runlock
[ ] add documentation to readme
[ ] add markdown to readme
[ ] push to github
*/

const (
	HOST    string = "localhost:8080"
	DB_PATH string = "database.json"
)

func main() {
	apiCfg := configure()

	CORSmux := apiCfg.createChiRouter()

	srv := &http.Server{
		Handler: CORSmux,
		Addr:    HOST,
	}

	log.Printf("Starting server at: %s", HOST)
	log.Fatal(srv.ListenAndServe())
}
