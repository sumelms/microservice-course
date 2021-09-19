package main

import (
	"context"
	"github.com/sumelms/microservice-course/internal/course"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/sumelms/microservice-course/pkg/config"
	database "github.com/sumelms/microservice-course/pkg/database/gorm"
	"golang.org/x/sync/errgroup"

	applogger "github.com/sumelms/microservice-course/pkg/logger"

	_ "github.com/lib/pq"
)

func main() {
	var (
		logger     log.Logger
		httpServer *http.Server
	)

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

	g.Go(func() error {
		httpLogger := log.With(logger, "component", "http")

		mux := http.NewServeMux()
		// Access Control and CORS
		http.Handle("/", accessControl(mux))

		// Initializing the services
		course.NewHTTPService(mux, db, httpLogger)

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
		httpServer.Shutdown(shutdownCtx) // nolint: errcheck
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
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
