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
	subscriptionRouter := NewSubscriptionRouter(s, opts...)

	r.PathPrefix("/courses").Handler(courseRouter)
	r.PathPrefix("/subscriptions").Handler(subscriptionRouter)
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

func NewSubscriptionRouter(s domain.ServiceInterface, opts ...kithttp.ServerOption) *mux.Router {
	r := mux.NewRouter().PathPrefix("/subscriptions").Subrouter().StrictSlash(true)

	listSubscriptionHandler := endpoints.NewListSubscriptionHandler(s, opts...)
	// listQueryParams := []string{"user_id", "{user_id}", "course_id", "{course_id}"}
	r.Handle("", listSubscriptionHandler).Methods(http.MethodGet)

	createSubscriptionHandler := endpoints.NewCreateSubscriptionHandler(s, opts...)
	r.Handle("", createSubscriptionHandler).Methods(http.MethodPost)

	findSubscriptionHandler := endpoints.NewFindSubscriptionHandler(s, opts...)
	r.Handle("/{uuid}", findSubscriptionHandler).Methods(http.MethodGet)

	deleteSubscriptionHandler := endpoints.NewDeleteSubscriptionHandler(s, opts...)
	r.Handle("/{uuid}", deleteSubscriptionHandler).Methods(http.MethodDelete)

	updateSubscriptionHandler := endpoints.NewUpdateSubscriptionHandler(s, opts...)
	r.Handle("/{uuid}", updateSubscriptionHandler).Methods(http.MethodPut)

	return r
}
