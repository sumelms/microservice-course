package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/go-kit/kit/endpoint"

	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/sumelms/microservice-course/internal/subscription/domain"
)

type findSubscriptionRequest struct {
	UUID uuid.UUID `json:"uuid"`
}

type findSubscriptionResponse struct {
	UUID       uuid.UUID  `json:"uuid"`
	UserID     uuid.UUID  `json:"user_id"`
	CourseID   uuid.UUID  `json:"course_id"`
	MatrixID   uuid.UUID  `json:"matrix_id"`
	ValidUntil *time.Time `json:"valid_until,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

func NewFindSubscriptionHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeFindSubscriptionEndpoint(s),
		decodeFindSubscriptionRequest,
		encodeFindSubscriptionResponse,
		opts...,
	)
}

func makeFindSubscriptionEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(findSubscriptionRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		sub, err := s.Subscription(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return &findSubscriptionResponse{
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

func decodeFindSubscriptionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	uid := uuid.MustParse(id)

	return findSubscriptionRequest{UUID: uid}, nil
}

func encodeFindSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
