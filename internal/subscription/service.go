package subscription

import (
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/subscription/database"
	"github.com/sumelms/microservice-course/internal/subscription/domain"
	courseService "github.com/sumelms/microservice-course/internal/subscription/services/course"
	"github.com/sumelms/microservice-course/internal/subscription/transport"
)

func NewHTTPService(router *mux.Router, db *sqlx.DB, logger log.Logger) error {
	service, err := NewService(db, logger)
	if err != nil {
		return err
	}
	transport.NewHTTPHandler(router, service, logger)
	return nil
}

func NewService(db *sqlx.DB, logger log.Logger) (*domain.Service, error) {
	repository, err := database.NewRepository(db)
	if err != nil {
		return nil, err
	}

	courseSvc, err := courseService.NewCourseService(db)
	if err != nil {
		logger.Log("subscription", "service", err, "msg", "unable to create service") // nolint: errcheck
		return nil, err
	}

	svc, err := domain.NewService(
		domain.WithRepository(repository),
		domain.WithLogger(logger),
		domain.WithCourseService(courseSvc),
	)
	if err != nil {
		return nil, err
	}
	return svc, nil
}
