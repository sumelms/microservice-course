package matrix

import (
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sumelms/microservice-course/internal/matrix/database"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
	"github.com/sumelms/microservice-course/internal/matrix/transport"
)

func NewService(db *sqlx.DB, logger log.Logger, course domain.CourseClient) (*domain.Service, error) {
	matrix, err := database.NewMatrixRepository(db)
	if err != nil {
		return nil, err
	}
	subject, err := database.NewSubjectRepository(db)
	if err != nil {
		return nil, err
	}

	service, err := domain.NewService(
		domain.WithLogger(logger),
		domain.WithMatrixRepository(matrix),
		domain.WithSubjectRepository(subject),
		domain.WithCourseClient(course))
	if err != nil {
		return nil, err
	}
	return service, nil
}

func NewHTTPService(router *mux.Router, service domain.ServiceInterface, logger log.Logger) error {
	transport.NewHTTPHandler(router, service, logger)
	return nil
}
