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

func NewHTTPService(router *mux.Router, db *sqlx.DB, logger log.Logger) {
	service := NewService(db, logger)
	transport.NewHTTPHandler(router, service, logger)
}

func NewService(db *sqlx.DB, logger log.Logger) *domain.Service {
	repository := &database.Repository{DB: db}
	courseSvc := clientService.NewCourseSvc(db, logger)
	return domain.NewService(repository, courseSvc, logger)
}
