package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/endpoint"
	"github.com/sumelms/microservice-course/pkg/validator"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

type updateMatrixRequest struct {
	UUID        string `json:"uuid" validate:"required"`
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description" validate:"required,max=255"`
}

type updateMatrixResponse struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewUpdateMatrixHandler(s domain.Service, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeUpdateMatrixEndpoint(s),
		decodeUpdateMatrixRequest,
		encodeUpdateMatrixResponse,
		opts...,
	)
}

func makeUpdateMatrixEndpoint(s domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(updateMatrixRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		var m domain.Matrix
		data, _ := json.Marshal(req)
		err := json.Unmarshal(data, &m)
		if err != nil {
			return nil, err
		}

		updated, err := s.UpdateMatrix(ctx, &m)
		if err != nil {
			return nil, err
		}

		return updateMatrixResponse{
			UUID:        updated.UUID,
			Title:       updated.Title,
			Description: updated.Description,
			CreatedAt:   updated.CreatedAt,
			UpdatedAt:   updated.UpdatedAt,
		}, nil
	}
}

func decodeUpdateMatrixRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	var req updateMatrixRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.UUID = id

	return req, nil
}

func encodeUpdateMatrixResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
