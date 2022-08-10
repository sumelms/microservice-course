package transport

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

	r.Handle("/courses", createCourseHandler).Methods(http.MethodPost)
	r.Handle("/courses", listCourseHandler).Methods(http.MethodGet)
	r.Handle("/courses/{uuid}", findCourseHandler).Methods(http.MethodGet)
	r.Handle("/courses/{uuid}", updateCourseHandler).Methods(http.MethodPut)
	r.Handle("/courses/{uuid}", deleteCourseHandler).Methods(http.MethodDelete)

	// Subscription handlers

	listSubscriptionHandler := endpoints.NewListSubscriptionHandler(s, opts...)
	createSubscriptionHandler := endpoints.NewCreateSubscriptionHandler(s, opts...)
	findSubscriptionHandler := endpoints.NewFindSubscriptionHandler(s, opts...)
	deleteSubscriptionHandler := endpoints.NewDeleteSubscriptionHandler(s, opts...)
	updateSubscriptionHandler := endpoints.NewUpdateSubscriptionHandler(s, opts...)

	r.Handle("/subscriptions", listSubscriptionHandler).Methods(http.MethodGet)
	r.Handle("/subscriptions", createSubscriptionHandler).Methods(http.MethodPost)
	r.Handle("/subscriptions/{uuid}", findSubscriptionHandler).Methods(http.MethodGet)
	r.Handle("/subscriptions/{uuid}", deleteSubscriptionHandler).Methods(http.MethodDelete)
	r.Handle("/subscriptions/{uuid}", updateSubscriptionHandler).Methods(http.MethodPut)
}
