package transport

import (
	"net/http"

	"github.com/sumelms/microservice-course/internal/course/endpoints"
	"github.com/sumelms/microservice-course/pkg/errors"

	"github.com/go-kit/kit/log"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/course/domain"
)

func NewHTTPHandler(r *mux.Router, s domain.ServiceInterface, logger log.Logger) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(errors.EncodeError),
	}

	listCourseHandler := endpoints.NewListCourseHandler(s, opts...)
	createCourseHandler := endpoints.NewCreateCourseHandler(s, opts...)
	findCourseHandler := endpoints.NewFindCourseHandler(s, opts...)
	updateCourseHandler := endpoints.NewUpdateCourseHandler(s, opts...)
	deleteCourseHandler := endpoints.NewDeleteCourseHandler(s, opts...)

	r.Handle("/courses", createCourseHandler).Methods(http.MethodPost)
	r.Handle("/courses", listCourseHandler).Methods(http.MethodGet)
	r.Handle("/courses/{uuid}", findCourseHandler).Methods(http.MethodGet)
	r.Handle("/courses/{uuid}", updateCourseHandler).Methods(http.MethodPut)
	r.Handle("/courses/{uuid}", deleteCourseHandler).Methods(http.MethodDelete)
}
