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

type updateSubscriptionRequest struct {
	UUID      uuid.UUID  `json:"uuid"       validate:"required"`
	Role      string     `json:"role"`
	ExpiresAt *time.Time `json:"expires_at"`
}

type updateSubscriptionResponse struct {
	Subscription *domain.Subscription `json:"subscription"`
}

func NewUpdateSubscriptionHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeUpdateSubscriptionEndpoint(s),
		decodeUpdateSubscriptionRequest,
		encodeUpdateSubscriptionResponse,
		opts...,
	)
}

func makeUpdateSubscriptionEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(updateSubscriptionRequest)
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

		if err := s.UpdateSubscription(ctx, &sub); err != nil {
			return nil, err
		}

		return updateSubscriptionResponse{
			Subscription: &sub,
		}, nil
	}
}

func decodeUpdateSubscriptionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	var req updateSubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	req.UUID = uuid.MustParse(id)

	return req, nil
}

func encodeUpdateSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
