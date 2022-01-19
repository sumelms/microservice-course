package subscription

import (
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/subscription/database"
	"github.com/sumelms/microservice-course/internal/subscription/domain"
	"github.com/sumelms/microservice-course/internal/subscription/transport"
)

func NewHTTPService(router *mux.Router, db *sqlx.DB, logger log.Logger) {
	repository := &database.Repository{DB: db}
	service := domain.NewService(repository, logger)

	transport.NewHTTPHandler(router, service, logger)
}
