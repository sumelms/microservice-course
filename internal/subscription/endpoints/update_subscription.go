package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/sumelms/microservice-course/pkg/validator"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/sumelms/microservice-course/internal/subscription/domain"
)

type updateSubscriptionRequest struct {
	ID         int        `json:"id" validate:"required"`
	UserID     string     `json:"user_id" validate:"required"`
	CourseID   string     `json:"course_id" validate:"required"`
	MatrixID   string     `json:"matrix_id" validate:"required"`
	ValidUntil *time.Time `json:"valid_until"`
}

type updateSubscriptionResponse struct {
	ID         uint       `json:"id"`
	UserID     string     `json:"user_id"`
	CourseID   string     `json:"course_id"`
	MatrixID   string     `json:"matrix_id"`
	ValidUntil *time.Time `json:"valid_until"`
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

		return updateSubscriptionResponse{
			ID:         sub.ID,
			UserID:     sub.UserID,
			CourseID:   sub.CourseID,
			MatrixID:   sub.MatrixID,
			ValidUntil: sub.ValidUntil,
			CreatedAt:  sub.CreatedAt,
			UpdatedAt:  sub.UpdatedAt,
		}, nil
	}
}

func decodeUpdateSubscriptionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	var req updateSubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	sid, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid argument conversion")
	}

	req.ID = sid

	return req, nil
}

func encodeUpdateSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
