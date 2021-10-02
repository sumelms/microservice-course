package transport

import (
	"net/http"

	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-course/internal/subscription/endpoints"
	"github.com/sumelms/microservice-course/pkg/errors"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/subscription/domain"
)

func NewHTTPHandler(r *mux.Router, s domain.Service, logger log.Logger) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(errors.EncodeError),
	}

	listSubscriptionHandler := endpoints.NewListSubscriptionHandler(s, opts...)
	findSubscriptionHandler := endpoints.NewFindSubscriptionHandler(s, opts...)
	createSubscriptionHandler := endpoints.NewCreateSubscriptionHandler(s, opts...)
	deleteSubscriptionHandler := endpoints.NewDeleteSubscriptionHandler(s, opts...)
	updateSubscriptionHandler := endpoints.NewUpdateSubscriptionHandler(s, opts...)
	findSubscriptionByCourseHandler := endpoints.NewFindSubscriptionByCourseHandler(s, opts...)
	findSubscriptionByUserHandler := endpoints.NewFindSubscriptionByUserHandler(s, opts...)

	r.Handle("/subscriptions", listSubscriptionHandler).Methods(http.MethodGet)
	r.Handle("/subscriptions/{uuid}", findSubscriptionHandler).Methods(http.MethodGet)
	r.Handle("/subscriptions", createSubscriptionHandler).Methods(http.MethodPost)
	r.Handle("/subscriptions/{uuid}", deleteSubscriptionHandler).Methods(http.MethodDelete)
	r.Handle("/subscriptions/{uuid}", updateSubscriptionHandler).Methods(http.MethodPut)
	r.Handle("/subscriptions/find-by-course-id/{uuid}", findSubscriptionByCourseHandler).Methods(http.MethodGet)
	r.Handle("/subscriptions/find-by-user-id/{uuid}", findSubscriptionByUserHandler).Methods(http.MethodGet)
}
