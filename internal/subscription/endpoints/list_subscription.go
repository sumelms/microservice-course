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
	CourseUUID uuid.UUID `json:"course_uuid"`
	UserUUID   uuid.UUID `json:"user_uuid"`
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
		if req.CourseUUID != uuid.Nil {
			filters.CourseUUID = req.CourseUUID
		}
		if req.UserUUID != uuid.Nil {
			filters.UserUUID = req.UserUUID
		}

		subscriptions, err := s.Subscriptions(ctx, filters)
		if err != nil {
			return nil, err
		}

		var list []findSubscriptionResponse
		for i := range subscriptions {
			sub := subscriptions[i]
			list = append(list, findSubscriptionResponse{
				UUID:       sub.UUID,
				UserUUID:   sub.UserUUID,
				CourseUUID: sub.CourseUUID,
				MatrixUUID: sub.MatrixUUID,
				Role:       sub.Role,
				ExpiresAt:  sub.ExpiresAt,
				CreatedAt:  sub.CreatedAt,
				UpdatedAt:  sub.UpdatedAt,
			})
		}

		return &listSubscriptionResponse{Subscriptions: list}, nil
	}
}

func decodeListSubscriptionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	courseUUID := r.FormValue("course_uuid")
	userUUID := r.FormValue("user_uuid")

	request := listSubscriptionRequest{}
	if len(courseUUID) > 0 {
		request.CourseUUID = uuid.MustParse(courseUUID)
	}
	if len(userUUID) > 0 {
		request.UserUUID = uuid.MustParse(userUUID)
	}

	return request, nil
}

func encodeListSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
