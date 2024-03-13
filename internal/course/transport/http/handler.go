package http

import (
	"net/http"

	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/internal/course/endpoints"
	"github.com/sumelms/microservice-course/pkg/errors"
)

func NewHTTPHandler(r *mux.Router, s domain.ServiceInterface, logger log.Logger) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(errors.EncodeError),
	}

	courseRouter := NewCourseRouter(s, opts...)

	r.PathPrefix("/courses").Handler(courseRouter)
}

func NewCourseRouter(s domain.ServiceInterface, opts ...kithttp.ServerOption) *mux.Router {
	r := mux.NewRouter().PathPrefix("/courses").Subrouter().StrictSlash(true)

	createCourseHandler := endpoints.NewCreateCourseHandler(s, opts...)
	r.Handle("", createCourseHandler).Methods(http.MethodPost)

	listCourseHandler := endpoints.NewListCourseHandler(s, opts...)
	r.Handle("", listCourseHandler).Methods(http.MethodGet)

	findCourseHandler := endpoints.NewFindCourseHandler(s, opts...)
	r.Handle("/{uuid}", findCourseHandler).Methods(http.MethodGet)

	updateCourseHandler := endpoints.NewUpdateCourseHandler(s, opts...)
	r.Handle("/{uuid}", updateCourseHandler).Methods(http.MethodPut)

	deleteCourseHandler := endpoints.NewDeleteCourseHandler(s, opts...)
	r.Handle("/{uuid}", deleteCourseHandler).Methods(http.MethodDelete)

	return r
}
