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

type removeSubjectRequest struct {
	MatrixID  uuid.UUID `json:"matrix_id" validate:"required"`
	SubjectID uuid.UUID `json:"subject_id" validate:"required"`
}

type removeSubjectResponse struct{}

func NewRemoveSubjectHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeRemoveSubjectEndpoint(s),
		decodeRemoveSubjectRequest,
		encodeRemoveSubjectResponse,
		opts...,
	)
}

func makeRemoveSubjectEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(removeSubjectRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		if err := s.RemoveSubject(ctx, req.MatrixID, req.SubjectID); err != nil {
			return nil, err
		}

		return removeSubjectResponse{}, nil
	}
}

func decodeRemoveSubjectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req removeSubjectRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeRemoveSubjectResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
