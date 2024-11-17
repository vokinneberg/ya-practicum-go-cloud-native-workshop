package handlers

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/vokinneberg/ya-practicum-go-cloud-native-workshop/internal/config"
	"github.com/vokinneberg/ya-practicum-go-cloud-native-workshop/internal/helpers"
	"github.com/vokinneberg/ya-practicum-go-cloud-native-workshop/internal/storage"
	"go.uber.org/zap"
)

type Storage interface {
	CreateShortURL(ctx context.Context, originalURL string, shortURL string) error
	GetOriginalURL(ctx context.Context, shortURL string) (string, error)
}

type Handler struct {
	config  *config.Config
	logger  *zap.Logger
	storage Storage
}

func New(storage Storage, config *config.Config, logger *zap.Logger) *Handler {
	return &Handler{storage: storage, config: config, logger: logger}
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	// Read plain text from the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}

	if len(body) == 0 {
		http.Error(w, "request body is empty", http.StatusBadRequest)
		return
	}

	// Create short URL
	shortURL := helpers.RandomString(6)
	if err := h.storage.CreateShortURL(r.Context(), string(body), shortURL); err != nil {
		if errors.Is(err, storage.ErrDuplicate) {
			http.Error(w, "short URL already exists", http.StatusConflict)
			return
		}
		http.Error(w, "failed to create short URL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(h.config.BaseURL + "/" + shortURL))
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) RedirectToOriginalURL(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	//Validate shortURL
	if id == "" {
		http.Error(w, "id is required", http.StatusNotFound)
		return
	}

	originalURL, err := h.storage.GetOriginalURL(r.Context(), id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			http.Error(w, "short URL not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to get original URL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", originalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
