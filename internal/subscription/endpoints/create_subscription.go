package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/sumelms/microservice-course/pkg/validator"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/sumelms/microservice-course/internal/subscription/domain"
)

type createSubscriptionRequest struct {
	UserID     uuid.UUID  `json:"user_id" validate:"required"`
	CourseID   uuid.UUID  `json:"course_id" validate:"required"`
	MatrixID   uuid.UUID  `json:"matrix_id" validate:"required"`
	ValidUntil *time.Time `json:"valid_until"`
}

type createSubscriptionResponse struct {
	UUID       uuid.UUID  `json:"uuid"`
	UserID     uuid.UUID  `json:"user_id"`
	CourseID   uuid.UUID  `json:"course_id"`
	MatrixID   uuid.UUID  `json:"matrix_id"`
	ValidUntil *time.Time `json:"valid_until"`
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
			UUID:       sub.UUID,
			UserID:     sub.UserID,
			CourseID:   sub.CourseID,
			MatrixID:   sub.MatrixID,
			ValidUntil: sub.ValidUntil,
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
