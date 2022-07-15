package subscription

import (
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/subscription/database"
	"github.com/sumelms/microservice-course/internal/subscription/domain"
	clientService "github.com/sumelms/microservice-course/internal/subscription/service"
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
	repository := &database.Repository{DB: db}
	courseSvc, err := clientService.NewCourseSvc(db, logger)
	if err != nil {
		return nil, err
	}
	return domain.NewService(repository, courseSvc, logger), nil
}
