package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sumelms/microservice-course/pkg/config"

	httptransport "github.com/sumelms/microservice-course/pkg/transport/http"

	"github.com/sumelms/microservice-course/pkg/logger"

	"github.com/go-kit/kit/log/level"
	_ "github.com/lib/pq"
)

func main() {
	// Logger
	logger := logger.NewLogger()

	cfg, err := loadConfig()
	if err != nil {
		_ = level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	_ = level.Info(logger).Log("msg", "service started")
	defer func() {
		_ = level.Info(logger).Log("msg", "service ended")
	}()

	// Database
	// db, err := database.Connect(cfg.Database)
	// if err != nil {
	// 	_ = level.Error(logger).Log("exit", err)
	// 	os.Exit(-1)
	// }

	ctx := context.Background()
	// repository := user.NewRepository(db, logger)
	// srv := userdomain.NewService(repository, logger)
	// endpoints := userendpoint.MakeEndpoints(srv)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// HTTP Server
	go func() {
		fmt.Println("HTTP Server Listening on", cfg.Server.HTTP.Host)
		httpServer := httptransport.NewHTTPServer(ctx)
		errs <- http.ListenAndServe(cfg.Server.HTTP.Host, httpServer)
	}()

	// gRPC Server
	go func() {
		// listener, err := net.Listen("tcp", cfg.Server.GRPC.Host)
		// if err != nil {
		// 	errs <- err
		// 	return
		// }

		fmt.Println("gRPC Server Listening on", cfg.Server.GRPC.Host)

		// handler := grpctransport.NewGRPCServer(ctx, endpoints)
		// grpcServer := grpc.NewServer()

		// protouser.RegisterUserServer(grpcServer, handler)
		// reflection.Register(grpcServer)

		// errs <- grpcServer.Serve(listener)
	}()

	_ = level.Error(logger).Log("exit", <-errs)
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
