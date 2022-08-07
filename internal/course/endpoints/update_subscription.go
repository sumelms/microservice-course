package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/sumelms/microservice-course/pkg/validator"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/sumelms/microservice-course/internal/course/domain"
)

type updateSubscriptionRequest struct {
	UUID       uuid.UUID  `json:"uuid" validate:"required"`
	UserID     uuid.UUID  `json:"user_id" validate:"required"`
	CourseID   uuid.UUID  `json:"course_id" validate:"required"`
	MatrixID   *uuid.UUID `json:"matrix_id"`
	ValidUntil *time.Time `json:"valid_until"`
}

type updateSubscriptionResponse struct {
	UUID       uuid.UUID  `json:"uuid"`
	UserID     uuid.UUID  `json:"user_id"`
	CourseID   uuid.UUID  `json:"course_id"`
	MatrixID   *uuid.UUID `json:"matrix_id,omitempty"`
	ValidUntil *time.Time `json:"valid_until,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
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
			UUID:       sub.UUID,
			UserID:     sub.UserID,
			CourseID:   sub.CourseID,
			MatrixID:   sub.MatrixID,
			ValidUntil: sub.ValidUntil,
			CreatedAt:  sub.CreatedAt,
			UpdatedAt:  sub.UpdatedAt,
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
