package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) metricsHandler(w http.ResponseWriter, req *http.Request) {
	respBody := "<html>\n\n<body>\n    <h1>Welcome, Chirpy Admin</h1>\n    <p>Chirpy has been visited %d times!</p>\n</body>\n\n</html>"
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(respBody, cfg.fileServerHits)))
	return
}
