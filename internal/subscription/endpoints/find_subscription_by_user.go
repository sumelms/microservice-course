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

type findSubscriptionByUserRequest struct {
	UserID string `json:"id"`
}

type findSubscriptionByUserResponse struct {
	Subscriptions []findSubscriptionResponse `json:"subscriptions"`
}

func NewFindSubscriptionByUserHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeFindSubscriptionByUserEndpoint(s),
		decodeFindSubscriptionByUserRequest,
		encodeFindSubscriptionByUserResponse,
		opts...,
	)
}

func makeFindSubscriptionByUserEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(findSubscriptionByUserRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		subscriptions, err := s.FindSubscriptionByUser(ctx, req.UserID)
		if err != nil {
			return nil, err
		}

		var list []findSubscriptionResponse
		for _, sub := range subscriptions {
			list = append(list, findSubscriptionResponse{
				ID:         sub.ID,
				UserID:     sub.UserID,
				CourseID:   sub.CourseID,
				MatrixID:   sub.MatrixID,
				ValidUntil: sub.ValidUntil,
				CreatedAt:  sub.CreatedAt,
				UpdatedAt:  sub.UpdatedAt,
			})
		}
		return &findSubscriptionByUserResponse{Subscriptions: list}, nil
	}
}

func decodeFindSubscriptionByUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	return findSubscriptionByUserRequest{UserID: id}, nil
}

func encodeFindSubscriptionByUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
