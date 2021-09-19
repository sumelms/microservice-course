package transport

import (
	"github.com/sumelms/microservice-course/internal/course/endpoints"
	"github.com/sumelms/microservice-course/pkg/errors"
	"net/http"

	"github.com/go-kit/kit/log"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/course/domain"
)

func NewHTTPHandler(s domain.Service, logger log.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(errors.EncodeError),
	}

	createCourseHandler := endpoints.NewCreateCourseHandler(s, opts...)
	getCourseHandler := endpoints.NewGetCourseHandler(s, opts...)
	updateCourseHandler := endpoints.NewUpdateCourseHandler(s, opts...)
	deleteCourseHandler := endpoints.NewDeleteCourseHandler(s, opts...)

	r := mux.NewRouter()

	r.Handle("/course", createCourseHandler).Methods(http.MethodPost)
	r.Handle("/course/{uuid}", getCourseHandler).Methods(http.MethodGet)
	r.Handle("/course/{uuid}", updateCourseHandler).Methods(http.MethodPut)
	r.Handle("/course/{uuid}", deleteCourseHandler).Methods(http.MethodDelete)

	return r
}
