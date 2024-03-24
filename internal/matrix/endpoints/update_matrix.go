package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
	"github.com/sumelms/microservice-course/pkg/validator"
)

type updateMatrixRequest struct {
	UUID        uuid.UUID `json:"uuid"        validate:"required"`
	Code        string    `json:"code"`
	Name        string    `json:"name"        validate:"required,max=100"`
	Description string    `json:"description" validate:"max=255"`
}

type updateMatrixResponse struct {
	Matrix *domain.Matrix `json:"matrix"`
}

func NewUpdateMatrixHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeUpdateMatrixEndpoint(s),
		decodeUpdateMatrixRequest,
		encodeUpdateMatrixResponse,
		opts...,
	)
}

//nolint:dupl
func makeUpdateMatrixEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(updateMatrixRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		var matrix domain.Matrix
		data, _ := json.Marshal(req)
		if err := json.Unmarshal(data, &matrix); err != nil {
			return nil, err
		}

		if err := s.UpdateMatrix(ctx, &matrix); err != nil {
			return nil, err
		}

		return updateMatrixResponse{
			Matrix: &matrix,
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

	req.UUID = uuid.MustParse(id)

	return req, nil
}

func encodeUpdateMatrixResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
