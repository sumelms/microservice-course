package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/sumelms/microservice-course/internal/subscription/domain"
)

type listSubscriptionRequest struct {
	CourseID uuid.UUID `json:"course_id"`
	UserID   uuid.UUID `json:"user_id"`
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

		filters := &domain.SubscriptionFilters{}
		if req.CourseID != uuid.Nil {
			filters.CourseID = req.CourseID
		}
		if req.UserID != uuid.Nil {
			filters.UserID = req.UserID
		}

		subscriptions, err := s.Subscriptions(ctx, filters)
		if err != nil {
			return nil, err
		}

		var list []findSubscriptionResponse
		for i := range subscriptions {
			sub := subscriptions[i]
			list = append(list, findSubscriptionResponse{
				UUID:      sub.UUID,
				UserID:    sub.UserID,
				CourseID:  sub.CourseID,
				MatrixID:  sub.MatrixID,
				Role:      sub.Role,
				ExpiresAt: sub.ExpiresAt,
				CreatedAt: sub.CreatedAt,
				UpdatedAt: sub.UpdatedAt,
			})
		}

		return &listSubscriptionResponse{Subscriptions: list}, nil
	}
}

func decodeListSubscriptionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	courseID := r.FormValue("course_id")
	userID := r.FormValue("user_id")

	request := listSubscriptionRequest{}
	if len(courseID) > 0 {
		request.CourseID = uuid.MustParse(courseID)
	}
	if len(userID) > 0 {
		request.UserID = uuid.MustParse(userID)
	}

	return request, nil
}

func encodeListSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
