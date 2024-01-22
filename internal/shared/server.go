package shared

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/pkg/config"
	"github.com/sumelms/microservice-course/pkg/logger"
)

type Server struct {
	*http.Server
	logger logger.Logger
	Config *config.HTTPServer
}

func NewServer(cfg *config.HTTPServer, router *mux.Router, log logger.Logger) (*Server, error) {
	if cfg == nil {
		return nil, fmt.Errorf("invalid server config")
	}
	if router == nil {
		return nil, fmt.Errorf("invalid server router")
	}

	return &Server{
		Server: &http.Server{
			Addr:    cfg.Host,
			Handler: router,
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		},
		logger: log,
		Config: cfg,
	}, nil
}

func (s *Server) Start() error {
	var err error

	s.logger.Log("msg", "starting HTTP Server", "host", s.Config.Host)

	if s.Config.UseHTTPS {
		s.logger.Log("msg", "SSL certificate enabled")

		certPath := s.Config.CertPath
		err = s.Server.ListenAndServeTLS(
			fmt.Sprintf("%s/server.crt", certPath),
			fmt.Sprintf("%s/server.key", certPath),
		)
	} else {
		s.logger.Log("msg", "SSL certificate disabled")
		err = s.Server.ListenAndServe()
	}

	if err != nil && err != http.ErrServerClosed {
		s.logger.Log("msg", "unable to start HTTP Server", "error", err)
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) {
	s.logger.Log("msg", "HTTP Server started to shutdown")

	err := s.Server.Shutdown(ctx)
	if err != nil {
		s.logger.Log("msg", "HTTP Server fails to shutdown", "error", err)
		return
	}

	s.logger.Log("msg", "HTTP Server shutdown successfully")
}
