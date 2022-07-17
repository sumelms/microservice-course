package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/go-kit/kit/endpoint"

	"github.com/sumelms/microservice-course/pkg/validator"

	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

type updateMatrixRequest struct {
	UUID        uuid.UUID `json:"uuid" validate:"required"`
	Code        string    `json:"code"`
	Name        string    `json:"name" validate:"required,max=100"`
	Description string    `json:"description" validate:"max=255"`
	CourseID    uuid.UUID `json:"course_id" validate:"required"`
}

type updateMatrixResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Code        string    `json:"code,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CourseID    uuid.UUID `json:"course_id"`
}

func NewUpdateMatrixHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeUpdateMatrixEndpoint(s),
		decodeUpdateMatrixRequest,
		encodeUpdateMatrixResponse,
		opts...,
	)
}

// nolint: dupl
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

		var m domain.Matrix
		data, _ := json.Marshal(req)
		if err := json.Unmarshal(data, &m); err != nil {
			return nil, err
		}

		if err := s.UpdateMatrix(ctx, &m); err != nil {
			return nil, err
		}

		return updateMatrixResponse{
			UUID:        m.UUID,
			Name:        m.Name,
			Description: m.Description,
			CreatedAt:   m.CreatedAt,
			UpdatedAt:   m.UpdatedAt,
			CourseID:    m.CourseID,
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
