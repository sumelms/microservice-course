package course

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/go-kit/kit/log"

	"github.com/sumelms/microservice-course/internal/course/database"
	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/internal/course/transport"
)

func NewHTTPService(router *mux.Router, db *sqlx.DB, logger log.Logger) {
	repository := &database.Repository{DB: db}
	service := domain.NewService(repository, logger)

	transport.NewHTTPHandler(router, service, logger)
}
