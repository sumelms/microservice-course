package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

type listMatrixRequest struct {
	CourseUUID uuid.UUID `json:"course_uuid,omitempty"`
}

type listMatrixResponse struct {
	Matrices []domain.Matrix `json:"matrices"`
}

func NewListMatrixHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeListMatrixEndpoint(s),
		decodeListMatrixRequest,
		encodeListMatrixResponse,
		opts...,
	)
}

func makeListMatrixEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(listMatrixRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		filters := &domain.MatrixFilters{}
		if req.CourseUUID != uuid.Nil {
			filters.CourseUUID = req.CourseUUID
		}

		matrices, err := s.Matrices(ctx, filters)
		if err != nil {
			return nil, err
		}

		return &listMatrixResponse{Matrices: matrices}, nil
	}
}

func decodeListMatrixRequest(_ context.Context, r *http.Request) (interface{}, error) {
	courseUUID := r.FormValue("course_uuid")

	request := listMatrixRequest{}
	if len(courseUUID) > 0 {
		request.CourseUUID = uuid.MustParse(courseUUID)
	}
	return request, nil
}

func encodeListMatrixResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
