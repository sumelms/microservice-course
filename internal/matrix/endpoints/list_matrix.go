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

type ListMatrixRequest struct {
	CourseUUID uuid.UUID `json:"course_uuid,omitempty"`
}

type ListMatrixResponse struct {
	Matrices []MatrixResponse `json:"matrices"`
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
		req, ok := request.(ListMatrixRequest)
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

		var list []MatrixResponse
		for i := range matrices {
			m := matrices[i]
			list = append(list, MatrixResponse{
				UUID:        m.UUID,
				Code:        m.Code,
				Name:        m.Name,
				Description: m.Description,
				CreatedAt:   m.CreatedAt,
				UpdatedAt:   m.UpdatedAt,
			})
		}

		return &ListMatrixResponse{Matrices: list}, nil
	}
}

func decodeListMatrixRequest(_ context.Context, r *http.Request) (interface{}, error) {
	courseUUID := r.FormValue("course_uuid")

	request := ListMatrixRequest{}
	if len(courseUUID) > 0 {
		request.CourseUUID = uuid.MustParse(courseUUID)
	}
	return request, nil
}

func encodeListMatrixResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
