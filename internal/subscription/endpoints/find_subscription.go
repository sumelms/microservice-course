package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/subscription/domain"
)

type FindSubscriptionRequest struct {
	UUID uuid.UUID `json:"uuid"`
}

type FindSubscriptionResponse struct {
	Subscription *SubscriptionResponse `json:"subscription"`
}

// NewFindSubscriptionHandler find subscription handler
// @Summary      Find subscription
// @Description  Find a new subscription
// @Tags         subscriptions
// @Produce      json
// @Param        uuid	  path      string  true  "Subscription UUID" Format(uuid)
// @Success      200      {object}  FindSubscriptionResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /subscriptions/{uuid} [get].
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
		req, ok := request.(FindSubscriptionRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		sub, err := s.Subscription(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return CreateSubscriptionResponse{
			Subscription: &SubscriptionResponse{
				UUID:       sub.UUID,
				UserUUID:   sub.UserUUID,
				CourseUUID: sub.CourseUUID,
				MatrixUUID: sub.MatrixUUID,
				Role:       sub.Role,
				ExpiresAt:  sub.ExpiresAt,
				CreatedAt:  sub.CreatedAt,
				UpdatedAt:  sub.UpdatedAt,
			},
		}, nil
	}
}

func decodeFindSubscriptionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	uid := uuid.MustParse(id)

	return FindSubscriptionRequest{UUID: uid}, nil
}

func encodeFindSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
