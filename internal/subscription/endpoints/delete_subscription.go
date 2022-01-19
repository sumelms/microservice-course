package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/sumelms/microservice-course/internal/subscription/domain"
)

type deleteSubscriptionRequest struct {
	ID int `json:"id" validate:"required"`
}

func NewDeleteSubscriptionHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeDeleteSubscriptionEndpoint(s),
		decodeDeleteSubscriptionRequest,
		encodeDeleteSubscriptionResponse,
		opts...,
	)
}

func makeDeleteSubscriptionEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(deleteSubscriptionRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		if err := s.DeleteSubscription(ctx, req.ID); err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func decodeDeleteSubscriptionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	sid, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid argument conversion")
	}

	return deleteSubscriptionRequest{ID: sid}, nil
}

func encodeDeleteSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
