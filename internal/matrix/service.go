package matrix

import (
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/matrix/database"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
	"github.com/sumelms/microservice-course/internal/matrix/transport"
)

func NewHTTPService(router *mux.Router, db *sqlx.DB, logger log.Logger) {
	repository := database.NewRepository(db)
	service := domain.NewService(repository, logger)

	transport.NewHTTPHandler(router, service, logger)
}
