package main

import (
	"net/http"
)

func methodNotAllowedHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}
