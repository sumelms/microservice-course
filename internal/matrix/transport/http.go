package transport

import (
	"net/http"

	"github.com/sumelms/microservice-course/internal/matrix/endpoints"
	"github.com/sumelms/microservice-course/pkg/errors"

	"github.com/go-kit/kit/log"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

func NewHTTPHandler(r *mux.Router, s domain.Service, logger log.Logger) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(errors.EncodeError),
	}

	listMatrixHandler := endpoints.NewListMatrixHandler(s, opts...)
	createMatrixHandler := endpoints.NewCreateMatrixHandler(s, opts...)
	findMatrixHandler := endpoints.NewFindMatrixHandler(s, opts...)
	updateMatrixHandler := endpoints.NewUpdateMatrixHandler(s, opts...)
	deleteMatrixHandler := endpoints.NewDeleteMatrixHandler(s, opts...)
	findMatrixByCourseHandler := endpoints.NewFindMatrixByCourse(s, opts...)

	r.Handle("/matrices", listMatrixHandler).Methods(http.MethodGet)
	r.Handle("/matrices", createMatrixHandler).Methods(http.MethodPost)
	r.Handle("/matrices/{uuid}", findMatrixHandler).Methods(http.MethodGet)
	r.Handle("/matrices/{uuid}", updateMatrixHandler).Methods(http.MethodPut)
	r.Handle("/matrices/{uuid}", deleteMatrixHandler).Methods(http.MethodDelete)
	r.Handle("/matrices/find-by-course-id/{uuid}", findMatrixByCourseHandler).Methods(http.MethodGet)
}
