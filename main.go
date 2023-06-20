package main

import (
	"log"
	"net/http"
)

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
