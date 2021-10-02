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

type findSubscriptionByCourseRequest struct {
	CourseID string `json:"id"`
}

type findSubscriptionByCourseResponse struct {
	Subscriptions []findSubscriptionResponse `json:"subscriptions"`
}

func NewFindSubscriptionByCourseHandler(s domain.Service, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeFindSubscriptionByCourseEndpoint(s),
		decodeFindSubscriptionByCourseRequest,
		encodeFindSubscriptionByCourseResponse,
		opts...,
	)
}

func makeFindSubscriptionByCourseEndpoint(s domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(findSubscriptionByCourseRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		subscriptions, err := s.FindSubscriptionByCourse(ctx, req.CourseID)
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
		return &findSubscriptionByCourseResponse{Subscriptions: list}, nil
	}
}

func decodeFindSubscriptionByCourseRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	return findSubscriptionByCourseRequest{CourseID: id}, nil
}

func encodeFindSubscriptionByCourseResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
