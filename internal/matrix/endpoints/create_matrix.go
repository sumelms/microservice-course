package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
	"github.com/sumelms/microservice-course/pkg/validator"
)

type createMatrixRequest struct {
	Code        string    `json:"code"        validate:"max=45"`
	Name        string    `json:"name"        validate:"required,max=100"`
	Description string    `json:"description" validate:"max=255"`
	CourseUUID  uuid.UUID `json:"course_uuid" validate:"required"`
}

type createMatrixResponse struct {
	Matrix *domain.Matrix `json:"matrix"`
}

func NewCreateMatrixHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeCreateMatrixEndpoint(s),
		decodeCreateMatrixRequest,
		encodeCreateMatrixResponse,
		opts...,
	)
}

//nolint:dupl
func makeCreateMatrixEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createMatrixRequest)
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

		if err := s.CreateMatrix(ctx, &matrix); err != nil {
			return nil, err
		}

		return createMatrixResponse{
			Matrix: &matrix,
		}, nil
	}
}

func decodeCreateMatrixRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createMatrixRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeCreateMatrixResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
