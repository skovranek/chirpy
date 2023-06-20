package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (cfg *apiConfig) createChiRouter() http.Handler {
	r := chi.NewRouter()
	fsHandler := cfg.metricsIncMiddleware(http.StripPrefix("/app", http.FileServer(http.Dir(cfg.root))))
	r.Handle("/app", fsHandler)
	r.Handle("/app/*", fsHandler)

	apiSubrouter := chi.NewRouter()
	apiSubrouter.Get("/healthz", readinessHandler)
	apiSubrouter.Post("/healthz", methodNotAllowedHandler)

	apiSubrouter.Post("/users", cfg.createUserHandler)
	apiSubrouter.Post("/login", cfg.loginHandler)
	apiSubrouter.Post("/refresh", cfg.refreshAccessTokenHandler)
	apiSubrouter.Post("/revoke", cfg.revokeRefreshTokenHandler)
	apiSubrouter.Put("/users", cfg.updateUserHandler)

	apiSubrouter.Post("/chirps", cfg.createChirpHandler)
	apiSubrouter.Get("/chirps", cfg.listChirpsHandler)
	apiSubrouter.Get("/chirps/{chirpID}", cfg.getChirpByIDHandler)
	apiSubrouter.Delete("/chirps/{chirpID}", cfg.deleteChirpHandler)

	apiSubrouter.Post("/polka/webhooks", cfg.polkaWebhooksHandler)
	r.Mount("/api", apiSubrouter)

	adminSubrouter := chi.NewRouter()
	adminSubrouter.Get("/metrics", cfg.metricsHandler)
	r.Mount("/admin", adminSubrouter)

	return CORSMiddleware(logMiddleware(r))
}
