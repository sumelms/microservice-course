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

type DeleteSubscriptionRequest struct {
	UUID   uuid.UUID `json:"uuid"   validate:"required"`
	Reason string    `json:"reason" validate:"required"`
}

type DeletedSubscriptionResponse struct {
	UUID      uuid.UUID `json:"uuid"`
	Reason    string    `json:"reason"`
	DeletedAt time.Time `json:"deleted_at"`
}

type DeleteSubscriptionResponse struct {
	Subscription *DeletedSubscriptionResponse `json:"subscription"`
}

// NewDeleteSubscriptionHandler deletes subscription handler
// @Summary      Delete subscription
// @Description  Delete a new subscription
// @Tags         subscriptions
// @Produce      json
// @Param        uuid	  path      string  true  "Subscription UUID" Format(uuid)
// @Success      200      {object}  DeleteSubscriptionResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /subscriptions/{uuid} [delete].
func NewDeleteSubscriptionHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeDeleteSubscriptionEndpoint(s),
		decodeDeleteSubscriptionRequest,
		encodeDeleteSubscriptionResponse,
		opts...,
	)
}

func makeDeleteSubscriptionEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(DeleteSubscriptionRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		var sub domain.DeletedSubscription
		data, _ := json.Marshal(req)
		if err := json.Unmarshal(data, &sub); err != nil {
			return nil, err
		}

		if err := s.DeleteSubscription(ctx, &sub); err != nil {
			return nil, err
		}

		return DeleteSubscriptionResponse{
			Subscription: &DeletedSubscriptionResponse{
				UUID:      sub.UUID,
				Reason:    sub.Reason,
				DeletedAt: sub.DeletedAt,
			},
		}, nil
	}
}

func decodeDeleteSubscriptionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	UUID, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	var req DeleteSubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	req.UUID = uuid.MustParse(UUID)

	return req, nil
}

func encodeDeleteSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
