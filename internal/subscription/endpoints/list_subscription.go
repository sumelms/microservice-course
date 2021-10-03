package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-course/internal/subscription/domain"
)

type listSubscriptionRequest struct {
	CourseID string `json:"course_id"`
	UserID   string `json:"user_id"`
}

type listSubscriptionResponse struct {
	Subscriptions []findSubscriptionResponse `json:"subscriptions"`
}

func NewListSubscriptionHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeListSubscriptionEndpoint(s),
		decodeListSubscriptionRequest,
		encodeListSubscriptionResponse,
		opts...,
	)
}

func makeListSubscriptionEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(listSubscriptionRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		filters := make(map[string]interface{})
		if len(req.CourseID) > 0 {
			filters["course_id"] = req.CourseID
		}
		if len(req.UserID) > 0 {
			filters["user_id"] = req.UserID
		}

		subscriptions, err := s.ListSubscription(ctx, filters)
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

		return &listSubscriptionResponse{Subscriptions: list}, nil
	}
}

func decodeListSubscriptionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	courseID := r.FormValue("course_id")
	userID := r.FormValue("user_id")
	return listSubscriptionRequest{
		CourseID: courseID,
		UserID:   userID,
	}, nil
}

func encodeListSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
