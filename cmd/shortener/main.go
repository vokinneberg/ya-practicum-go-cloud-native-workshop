package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/vokinneberg/ya-practicum-go-cloud-native-workshop/internal/config"
	"github.com/vokinneberg/ya-practicum-go-cloud-native-workshop/internal/http/handlers"
	"github.com/vokinneberg/ya-practicum-go-cloud-native-workshop/internal/logging"
	"github.com/vokinneberg/ya-practicum-go-cloud-native-workshop/internal/storage"
	"go.uber.org/zap"

	internalHttp "github.com/vokinneberg/ya-practicum-go-cloud-native-workshop/internal/http"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to create config: %v", err)
	}
	logger := logging.New(cfg, os.Stdout)

	db, err := sql.Open("postgres", cfg.DatabaseDSN)
	if err != nil {
		logger.Error("Failed to open database", zap.Error(err))
		os.Exit(1)
	}
	defer db.Close()

	// Run migrations
	if cfg.RunMigrations {
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			logger.Error("Failed to create driver", zap.Error(err))
			os.Exit(1)
		}
		m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "postgres", driver)
		if err != nil {
			logger.Error("Failed to create migrations", zap.Error(err))
			os.Exit(1)
		}
		if err := m.Up(); err != nil {
			if !errors.Is(err, migrate.ErrNoChange) {
				logger.Error("Failed to run migrations", zap.Error(err))
				os.Exit(1)
			}
		}
	}

	// Create storage
	storage, err := storage.NewPostgresStorage(db)
	if err != nil {
		logger.Error("Failed to create storage", zap.Error(err))
		os.Exit(1)
	}

	// Start HTTP server
	router := internalHttp.NewRouter(http.NewServeMux(), logger)
	handler := handlers.New(storage, cfg, logger)
	mux := router.RegisterRoutes(handler)

	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: mux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		logger.Info("Starting server", zap.String("address", cfg.ServerAddress))
		if err := server.ListenAndServe(); err != nil {
			logger.Error("Failed to start server", zap.Error(err))
		}
	}()

	<-stop
	logger.Info("Shutting down server...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Failed to shutdown server", zap.Error(err))
		os.Exit(1)
	}

	logger.Info("Server shutdown")
}
