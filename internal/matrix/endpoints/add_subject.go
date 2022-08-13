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

type addSubjectRequest struct {
	MatrixID  uuid.UUID `json:"matrix_id" validate:"required"`
	SubjectID uuid.UUID `json:"subject_id" validate:"required"`
}

type addSubjectResponse struct{}

func NewAddSubjectHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeAddSubjectEndpoint(s),
		decodeAddSubjectRequest,
		encodeAddSubjectResponse,
		opts...,
	)
}

//nolint:dupl
func makeAddSubjectEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(addSubjectRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		ms := &domain.MatrixSubject{MatrixID: req.MatrixID, SubjectID: req.SubjectID}
		if err := s.AddSubject(ctx, ms); err != nil {
			return nil, err
		}

		return addSubjectResponse{}, nil
	}
}

func decodeAddSubjectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req addSubjectRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeAddSubjectResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
