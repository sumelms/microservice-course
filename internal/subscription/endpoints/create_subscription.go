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
	"github.com/sumelms/microservice-course/internal/subscription/domain"
	"github.com/sumelms/microservice-course/pkg/validator"
)

type createSubscriptionRequest struct {
	UserUUID   uuid.UUID  `json:"user_uuid"   validate:"required"`
	CourseUUID uuid.UUID  `json:"course_uuid" validate:"required"`
	MatrixUUID *uuid.UUID `json:"matrix_uuid"`
	Role       string     `json:"role"        validate:"required"`
	ExpiresAt  *time.Time `json:"expires_at"`
}

type createSubscriptionResponse struct {
	Subscription *domain.Subscription `json:"subscription"`
}

func NewCreateSubscriptionHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeCreateSubscriptionEndpoint(s),
		decodeCreateSubscriptionRequest,
		encodeCreateSubscriptionResponse,
		opts...,
	)
}

func makeCreateSubscriptionEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createSubscriptionRequest)
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

		if err := s.CreateSubscription(ctx, &sub); err != nil {
			return nil, err
		}
		return createSubscriptionResponse{
			Subscription: &sub,
		}, nil
	}
}

func decodeCreateSubscriptionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createSubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func encodeCreateSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
