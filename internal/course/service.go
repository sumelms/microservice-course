package course

import (
	"net/http"

	"github.com/go-kit/kit/log"

	"github.com/jinzhu/gorm"
	"github.com/sumelms/microservice-course/internal/course/database"
	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/internal/course/transport"
)

func NewHTTPService(mux *http.ServeMux, db *gorm.DB, logger log.Logger) {
	repository := database.NewRepository(db, logger)
	service := domain.NewService(repository, logger)

	mux.Handle("/", transport.NewHTTPHandler(service, logger))
}
