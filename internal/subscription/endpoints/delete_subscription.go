package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-course/internal/subscription/domain"
)

type deleteSubscriptionRequest struct {
	ID string `json:"id" validate:"required"`
}

func NewDeleteSubscriptionHandler(s domain.Service, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeDeleteSubscriptionEdnpoint(s),
		decodeDeleteSubscriptionRequest,
		encodeDeleteSubscriptionResponse,
		opts...,
	)
}

func makeDeleteSubscriptionEdnpoint(s domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(deleteSubscriptionRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		err := s.DeleteSubscription(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func decodeDeleteSubscriptionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	return deleteSubscriptionRequest{ID: id}, nil
}

func encodeDeleteSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
