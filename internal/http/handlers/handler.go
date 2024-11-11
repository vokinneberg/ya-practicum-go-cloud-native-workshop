package handlers

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

type Storage interface {
	CreateShortURL(ctx context.Context, originalURL string) (string, error)
	GetOriginalURL(ctx context.Context, shortURL string) (string, error)
}

type Handler struct {
	storage Storage
	logger  *zap.Logger
}

func New(storage Storage, logger *zap.Logger) *Handler {
	return &Handler{storage: storage, logger: logger}
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("shortening URL")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) RedirectToOriginalURL(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("redirecting to original URL")
	w.WriteHeader(http.StatusOK)
}
