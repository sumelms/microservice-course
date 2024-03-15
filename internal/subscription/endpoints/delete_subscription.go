package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/subscription/domain"
	"github.com/sumelms/microservice-course/pkg/validator"
)

type deleteSubscriptionRequest struct {
	UUID   uuid.UUID `json:"uuid"   validate:"required"`
	Reason string    `json:"reason" validate:"required"`
}

type deleteSubscriptionResponse struct {
	UUID      uuid.UUID `json:"uuid"`
	Reason    string    `json:"reason"`
	DeletedAt time.Time `json:"deleted_at"`
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

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		var sub domain.Subscription
		data, _ := json.Marshal(req)
		if err := json.Unmarshal(data, &sub); err != nil {
			return nil, err
		}

		if err := s.DeleteSubscription(ctx, &sub); err != nil {
			return nil, err
		}

		return deleteSubscriptionResponse{
			UUID:      sub.UUID,
			Reason:    *sub.Reason,
			DeletedAt: *sub.DeletedAt,
		}, nil
	}
}

func decodeDeleteSubscriptionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	UUID, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	var req deleteSubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	req.UUID = uuid.MustParse(UUID)

	return req, nil
}

func encodeDeleteSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
