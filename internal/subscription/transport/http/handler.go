package http

import (
	"net/http"

	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/subscription/domain"
	"github.com/sumelms/microservice-course/internal/subscription/endpoints"
	"github.com/sumelms/microservice-course/pkg/errors"
)

func NewHTTPHandler(r *mux.Router, s domain.ServiceInterface, logger log.Logger) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(errors.EncodeError),
	}

	subscriptionRouter := NewSubscriptionRouter(s, opts...)

	r.PathPrefix("/subscriptions").Handler(subscriptionRouter)
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
