package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/sumelms/microservice-course/internal/subscription/domain"
)

type ListSubscriptionsRequest struct {
	CourseUUID uuid.UUID `json:"course_uuid"`
	UserUUID   uuid.UUID `json:"user_uuid"`
}

type SubscriptionCourseResponse struct {
	UUID uuid.UUID `db:"uuid" json:"uuid"`
	Code string    `db:"code" json:"code"`
	Name string    `db:"name" json:"name"`
}

type SubscriptionMatrixResponse struct {
	UUID *uuid.UUID `db:"uuid" json:"uuid,omitempty"`
	Code *string    `db:"code" json:"code,omitempty"`
	Name *string    `db:"name" json:"name,omitempty"`
}

type SubscriptionsResponse struct {
	UUID       uuid.UUID                   `json:"uuid"`
	UserUUID   uuid.UUID                   `json:"user_uuid"`
	Course     *SubscriptionCourseResponse `json:"course,omitempty"`
	CourseUUID *uuid.UUID                  `json:"course_uuid,omitempty"`
	Matrix     *SubscriptionMatrixResponse `json:"matrix,omitempty"`
	MatrixUUID *uuid.UUID                  `json:"matrix_uuid,omitempty"`
	Role       string                      `json:"role"`
	ExpiresAt  *time.Time                  `json:"expires_at,omitempty"`
	CreatedAt  time.Time                   `json:"created_at"`
	UpdatedAt  time.Time                   `json:"updated_at"`
}

type ListSubscriptionsResponse struct {
	Subscriptions []SubscriptionsResponse `json:"subscriptions"`
}

// NewListSubscriptionsHandler list subscriptions handler
// @Summary      List subscriptions
// @Description  List a new subscriptions
// @Tags         subscriptions
// @Produce      json
// @Param        course_uuid    query     string  false  "course search by uuid"  Format(uuid)
// @Param        user_uuid    	query     string  false  "user search by uuid"  Format(uuid)
// @Success      200      {object}  ListSubscriptionsResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /subscriptions [get].
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

		var list []SubscriptionsResponse
		for i := range subscriptions {
			sub := subscriptions[i]
			list = append(list, SubscriptionsResponse{
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
		parsedCourseUUID, err := uuid.Parse(courseUUID)
		if err != nil {
			return nil, fmt.Errorf("invalid course UUID: %v", err)
		}
		request.CourseUUID = parsedCourseUUID
	}
	if len(userUUID) > 0 {
		parsedUserUUID, err := uuid.Parse(userUUID)
		if err != nil {
			return nil, fmt.Errorf("invalid user UUID: %v", err)
		}
		request.UserUUID = parsedUserUUID
	}

	return request, nil
}

func encodeListSubscriptionsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
