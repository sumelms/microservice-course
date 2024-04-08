package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/subscription/domain"
	"github.com/sumelms/microservice-course/pkg/validator"
)

type UpdateSubscriptionRequest struct {
	UUID      uuid.UUID  `json:"uuid"       validate:"required"`
	Role      string     `json:"role"       validate:"required"`
	ExpiresAt *time.Time `json:"expires_at"`
}

type UpdateSubscriptionResponse struct {
	Subscription *SubscriptionResponse `json:"subscription"`
}

// NewUpdateSubscriptionHandler updates new subscription handler
// @Summary      Update subscription
// @Description  Update a subscription
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        uuid	  		  path      string  true  "Subscription UUID" Format(uuid)
// @Param        subscription	  body		UpdateSubscriptionRequest		true	"Update Subscription"
// @Success      200      {object}  UpdateSubscriptionResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /subscriptions/{uuid} [put].
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
		req, ok := request.(UpdateSubscriptionRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		var subscription domain.Subscription
		data, _ := json.Marshal(req)
		if err := json.Unmarshal(data, &subscription); err != nil {
			return nil, err
		}

		sub, err := s.UpdateSubscription(ctx, &subscription)
		if err != nil {
			return nil, err
		}

		return UpdateSubscriptionResponse{
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

func decodeUpdateSubscriptionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	var req UpdateSubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	req.UUID = uuid.MustParse(id)

	return req, nil
}

func encodeUpdateSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
