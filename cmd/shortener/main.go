package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vokinneberg/ya-practicum-go-cloud-native-workshop/internal/config"
	"github.com/vokinneberg/ya-practicum-go-cloud-native-workshop/internal/http/handlers"
	"github.com/vokinneberg/ya-practicum-go-cloud-native-workshop/internal/logging"
	"go.uber.org/zap"

	internalHttp "github.com/vokinneberg/ya-practicum-go-cloud-native-workshop/internal/http"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to create config: %v", err)
	}

	logger := logging.New(cfg.AppName, cfg.LogLevel, cfg.Environment, cfg.Version, os.Stdout)

	handler := handlers.New(nil, logger)
	router := internalHttp.NewRouter(http.NewServeMux(), logger)
	mux := router.RegisterHandlers(handler)

	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: mux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.Info("starting server", zap.String("address", cfg.ServerAddress))
		if err := server.ListenAndServe(); err != nil {
			logger.Error("failed to start server", zap.Error(err))
			os.Exit(1)
		}
	}()

	<-stop
	logger.Info("shutting down server")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	logger.Info("shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("failed to shutdown server", zap.Error(err))
		os.Exit(1)
	}

	logger.Info("server shutdown")
}
