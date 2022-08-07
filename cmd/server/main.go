package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"github.com/sumelms/microservice-course/internal/matrix"

	"github.com/sumelms/microservice-course/internal/course"

	"github.com/go-kit/log"
	"golang.org/x/sync/errgroup"

	"github.com/sumelms/microservice-course/pkg/config"
	database "github.com/sumelms/microservice-course/pkg/database/postgres"

	applogger "github.com/sumelms/microservice-course/pkg/logger"

	_ "github.com/lib/pq"
)

var (
	logger     log.Logger
	httpServer *http.Server
)

func main() {
	// Logger
	logger = applogger.NewLogger()
	logger.Log("msg", "service started") // nolint: errcheck

	// Configuration
	cfg, err := loadConfig()
	if err != nil {
		logger.Log("exit", err) // nolint: errcheck
		os.Exit(-1)
	}

	// Database
	db, err := database.Connect(cfg.Database)
	if err != nil {
		logger.Log("msg", "database error", err) // nolint: errcheck
		os.Exit(1)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	// Initialize the domain services
	svcLogger := log.With(logger, "component", "service")

	courseSvc, err := course.NewService(db, svcLogger)
	if err != nil {
		logger.Log("msg", "unable to start course service", err) // nolint: errcheck
		os.Exit(1)
	}
	matrixSvc, err := matrix.NewService(db, svcLogger)
	if err != nil {
		logger.Log("msg", "unable to start matrix service", err) // nolint: errcheck
		os.Exit(1)
	}

	g.Go(func() error {
		// Initialize the router
		router := mux.NewRouter()

		// Initializing the HTTP Services
		httpLogger := log.With(logger, "component", "http")

		if err := course.NewHTTPService(router, courseSvc, httpLogger); err != nil {
			logger.Log("msg", "unable to start a service: course", "error", err) // nolint: errcheck
			return err
		}
		if err := matrix.NewHTTPService(router, matrixSvc, httpLogger); err != nil {
			logger.Log("msg", "unable to start a service: matrix", "error", err) // nolint: errcheck
			return err
		}

		// Handle the mux & router
		srv := http.NewServeMux()
		srv.Handle("/", router)

		// Middlewares
		http.Handle("/", accessControl(srv))

		logger.Log("transport", "http", "address", cfg.Server.HTTP.Host, "msg", "listening") // nolint: errcheck

		httpServer = &http.Server{
			Addr:         cfg.Server.HTTP.Host,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	logger.Log("msg", "received shutdown signal") // nolint: errcheck

	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if httpServer != nil {
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			logger.Log("msg", "server wasn't gracefully shutdown") // nolint: errcheck
			defer os.Exit(2)
		}
	}

	if err := g.Wait(); err != nil {
		logger.Log("msg", "server returning an error", "error", err) // nolint: errcheck
		defer os.Exit(2)
	}

	logger.Log("msg", "service ended") // nolint: errcheck
}

func loadConfig() (*config.Config, error) {
	// Configuration
	configPath := os.Getenv("SUMELMS_CONFIG_PATH")
	if configPath == "" {
		configPath = "./config.yml"
	}

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
