package course

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/go-kit/log"

	"github.com/sumelms/microservice-course/internal/course/database"
	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/internal/course/transport"
)

func NewHTTPService(router *mux.Router, db *sqlx.DB, logger log.Logger) error {
	service, err := NewService(db, logger)
	if err != nil {
		return err
	}
	transport.NewHTTPHandler(router, service, logger)
	return nil
}

func NewService(db *sqlx.DB, logger log.Logger) (*domain.service, error) {
	repository, err := database.NewCourseRepository(db)
	if err != nil {
		return nil, err
	}
	return domain.NewService(repository, logger), nil
}
