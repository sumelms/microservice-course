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

type ListSubscriptionsRequest struct {
	CourseUUID uuid.UUID `json:"course_uuid"`
	UserUUID   uuid.UUID `json:"user_uuid"`
}

type ListSubscriptionsResponse struct {
	Subscriptions []SubscriptionResponse `json:"subscriptions"`
}

func NewListSubscriptionsHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeListSubscriptionsEndpoint(s),
		decodeListSubscriptionsRequest,
		encodeListSubscriptionsResponse,
		opts...,
	)
}

func SerializeCourse(subscription domain.Subscription) *SubscriptionCourseResponse {
	if subscription.Course != nil {
		return &SubscriptionCourseResponse{
			UUID: subscription.Course.UUID,
			Code: subscription.Course.Code,
			Name: subscription.Course.Name,
		}
	}

	return nil
}

func SerializeMatrix(subscription domain.Subscription) *SubscriptionMatrixResponse {
	if subscription.Matrix != nil {
		return &SubscriptionMatrixResponse{
			UUID: subscription.Matrix.UUID,
			Code: subscription.Matrix.Code,
			Name: subscription.Matrix.Name,
		}
	}

	return nil
}

func makeListSubscriptionsEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(ListSubscriptionsRequest)
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

		var list []SubscriptionResponse
		for i := range subscriptions {
			sub := subscriptions[i]
			list = append(list, SubscriptionResponse{
				UUID:       sub.UUID,
				UserUUID:   sub.UserUUID,
				CourseUUID: sub.CourseUUID,
				Course:     SerializeCourse(sub),
				MatrixUUID: sub.MatrixUUID,
				Matrix:     SerializeMatrix(sub),
				Role:       sub.Role,
				ExpiresAt:  sub.ExpiresAt,
				CreatedAt:  sub.CreatedAt,
				UpdatedAt:  sub.UpdatedAt,
			})
		}
		return &ListSubscriptionsResponse{Subscriptions: list}, nil
	}
}

func decodeListSubscriptionsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	courseUUID := r.FormValue("course_uuid")
	userUUID := r.FormValue("user_uuid")

	request := ListSubscriptionsRequest{}
	if len(courseUUID) > 0 {
		request.CourseUUID = uuid.MustParse(courseUUID)
	}
	if len(userUUID) > 0 {
		request.UserUUID = uuid.MustParse(userUUID)
	}

	return request, nil
}

func encodeListSubscriptionsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
