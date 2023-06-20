package main

import (
	"log"
	"net/http"
)

/*
TODO:
[x] fix lack of auth headers causing index panic when len called on missing param string
[x] get rid of get user by ID, redundant since maps return key
[x] refactor
[x] extract
[x] extract server
[x] extract auth
[x] extract in db
[x] remove JWT and bcrypt from database
[x] errors fmt
[x] change constants to env_vars

[x] add mux Rlock defer Runlock
[ ] add documentation to readme
[ ] add markdown to readme
[ ] push to github
*/

func main() {
	apiCfg := configure()

	CORSmux := apiCfg.createChiRouter()

	srv := &http.Server{
		Handler: CORSmux,
		Addr:    ":" + apiCfg.port,
	}

	log.Printf("Starting server at port #%s", apiCfg.port)
	log.Fatal(srv.ListenAndServe())
}
