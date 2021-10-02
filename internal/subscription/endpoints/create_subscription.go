package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sumelms/microservice-course/pkg/validator"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-course/internal/subscription/domain"
)

type createSubscriptionRequest struct {
	UserID     string     `json:"user_id" validate:"required"`
	CourseID   string     `json:"course_id" validate:"required"`
	MatrixID   string     `json:"matrix_id" validate:"required"`
	ValidUntil *time.Time `json:"valid_until"`
}

type createSubscriptionResponse struct {
	UserID     string     `json:"user_id"`
	CourseID   string     `json:"course_id"`
	MatrixID   string     `json:"matrix_id"`
	ValidUntil *time.Time `json:"valid_until"`
}

func NewCreateSubscriptionHandler(s domain.Service, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeCreateSubscriptionEndpoint(s),
		decodeCreateSubscriptionRequest,
		encodeCreateSubscriptionResponse,
		opts...,
	)
}

func makeCreateSubscriptionEndpoint(s domain.Service) endpoint.Endpoint {
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
		err := json.Unmarshal(data, &sub)
		if err != nil {
			return nil, err
		}

		created, err := s.CreateSubscription(ctx, &sub)
		if err != nil {
			return nil, err
		}

		return createSubscriptionResponse{
			UserID:     created.UserID,
			CourseID:   created.CourseID,
			MatrixID:   created.MatrixID,
			ValidUntil: created.ValidUntil,
		}, nil
	}
}

func decodeCreateSubscriptionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req createSubscriptionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeCreateSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
