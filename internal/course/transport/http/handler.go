package http

import (
	"net/http"

	"github.com/sumelms/microservice-course/internal/course/endpoints"
	"github.com/sumelms/microservice-course/pkg/errors"

	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"

	"github.com/sumelms/microservice-course/internal/course/domain"
)

func NewHTTPHandler(r *mux.Router, s domain.ServiceInterface, logger log.Logger) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(errors.EncodeError),
	}

	// Course handlers
	listCourseHandler := endpoints.NewListCourseHandler(s, opts...)
	createCourseHandler := endpoints.NewCreateCourseHandler(s, opts...)
	findCourseHandler := endpoints.NewFindCourseHandler(s, opts...)
	updateCourseHandler := endpoints.NewUpdateCourseHandler(s, opts...)
	deleteCourseHandler := endpoints.NewDeleteCourseHandler(s, opts...)

	cr := r.PathPrefix("/courses").Subrouter().StrictSlash(true)

	cr.Handle("/", createCourseHandler).Methods(http.MethodPost)
	cr.Handle("/", listCourseHandler).Methods(http.MethodGet)
	cr.Handle("/{uuid}", findCourseHandler).Methods(http.MethodGet)
	cr.Handle("/{uuid}", updateCourseHandler).Methods(http.MethodPut)
	cr.Handle("/{uuid}", deleteCourseHandler).Methods(http.MethodDelete)

	// Subscription handlers
	listSubscriptionHandler := endpoints.NewListSubscriptionHandler(s, opts...)
	createSubscriptionHandler := endpoints.NewCreateSubscriptionHandler(s, opts...)
	findSubscriptionHandler := endpoints.NewFindSubscriptionHandler(s, opts...)
	deleteSubscriptionHandler := endpoints.NewDeleteSubscriptionHandler(s, opts...)
	updateSubscriptionHandler := endpoints.NewUpdateSubscriptionHandler(s, opts...)

	sr := r.PathPrefix("/subscriptions").Subrouter().StrictSlash(true)

	sr.Handle("/", listSubscriptionHandler).Methods(http.MethodGet)
	sr.Handle("/", createSubscriptionHandler).Methods(http.MethodPost)
	sr.Handle("/{uuid}", findSubscriptionHandler).Methods(http.MethodGet)
	sr.Handle("/{uuid}", deleteSubscriptionHandler).Methods(http.MethodDelete)
	sr.Handle("/{uuid}", updateSubscriptionHandler).Methods(http.MethodPut)
}
