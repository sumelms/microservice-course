package subject

import (
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/subject/database"
	"github.com/sumelms/microservice-course/internal/subject/domain"
	"github.com/sumelms/microservice-course/internal/subject/transport"
)

func NewHTTPService(router *mux.Router, db *sqlx.DB, logger log.Logger) {
	service := NewService(db, logger)
	transport.NewHTTPHandler(router, service, logger)
}

func NewService(db *sqlx.DB, logger log.Logger) *domain.Service {
	repository := &database.Repository{DB: db}
	return domain.NewService(repository, logger)
}
