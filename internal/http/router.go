package http

import (
	"net/http"

	"github.com/go-pkgz/routegroup"
	"github.com/vokinneberg/ya-practicum-go-cloud-native-workshop/internal/http/handlers"
	"github.com/vokinneberg/ya-practicum-go-cloud-native-workshop/internal/http/middleware"
	"go.uber.org/zap"
)

type Router struct {
	mux    *http.ServeMux
	logger *zap.Logger
}

// NewRouter creates a new router
func NewRouter(mux *http.ServeMux, logger *zap.Logger) *Router {
	return &Router{
		mux:    http.NewServeMux(),
		logger: logger,
	}
}

// RegisterHandlers registers the handlers for the router
func (rr *Router) RegisterRoutes(handler *handlers.Handler) http.Handler {
	group := routegroup.New(rr.mux)
	group.Use(middleware.Logging(rr.logger))

	group.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("pong")); err != nil {
			rr.logger.Error("failed to write response", zap.Error(err))
		}
		w.WriteHeader(http.StatusOK)
	})

	// Liveness probe endpoint
	group.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Readiness probe endpoint
	group.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
		// Add logic to check if the server is ready
		// For now, we assume it's always ready
		w.WriteHeader(http.StatusOK)
	})

	// Create a new short URL
	group.HandleFunc("POST /", handler.ShortenURL)

	// Redirect to the original URL
	group.HandleFunc("GET /{id}", handler.RedirectToOriginalURL)

	return group
}
