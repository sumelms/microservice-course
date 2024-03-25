package subscription

import (
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sumelms/microservice-course/internal/subscription/database"
	"github.com/sumelms/microservice-course/internal/subscription/domain"
	"github.com/sumelms/microservice-course/internal/subscription/transport/http"
)

func NewService(db *sqlx.DB,
	logger log.Logger,
	course domain.CourseClient,
	matrix domain.MatrixClient,
) (*domain.Service, error) {
	subscription, err := database.NewSubscriptionRepository(db)
	if err != nil {
		return nil, err
	}

	service, err := domain.NewService(
		domain.WithLogger(logger),
		domain.WithSubscriptionRepository(subscription),
		domain.WithCourseClient(course),
		domain.WithMatrixClient(matrix))
	if err != nil {
		return nil, err
	}
	return service, nil
}

func NewHTTPService(router *mux.Router, service domain.ServiceInterface, logger log.Logger) error {
	http.NewHTTPHandler(router, service, logger)
	return nil
}
