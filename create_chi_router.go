package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (cfg *apiConfig) createChiRouter() http.Handler {
	r := chi.NewRouter()
	r.Mount("/", cfg.middlewareMetricsInc(middlewareLog(http.FileServer(http.Dir(".")))))

	r.Get("/admin/metrics", cfg.metricsHandler)

	apiSubrouter := chi.NewRouter()
	apiSubrouter.Get("/healthz", healthzHandler)
	apiSubrouter.Post("/healthz", methodNotAllowedHandler)
	apiSubrouter.Post("/chirps", cfg.createChirpHandler)
	apiSubrouter.Get("/chirps", cfg.getChirpsHandler)
	apiSubrouter.Get("/chirps/{chirpID}", cfg.getChirpByIDHandler)
	apiSubrouter.Delete("/chirps/{chirpID}", cfg.deleteChirpHandler)
	apiSubrouter.Post("/users", cfg.createUserHandler)
	apiSubrouter.Put("/users", cfg.updateUserHandler)
	apiSubrouter.Post("/login", cfg.loginHandler)
	apiSubrouter.Post("/refresh", cfg.refreshHandler)
	apiSubrouter.Post("/revoke", cfg.revokeHandler)
	apiSubrouter.Post("/polka/webhooks", cfg.polkaWebhooksHandler)
	r.Mount("/api", middlewareLog(apiSubrouter))

	return middlewareCORS(r)
}
