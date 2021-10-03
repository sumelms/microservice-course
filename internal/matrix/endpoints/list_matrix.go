package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

type listMatrixRequest struct {
	CourseID string `json:"course_id,omitempty"`
}

type listMatrixResponse struct {
	Matrices []findMatrixResponse `json:"matrices"`
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

		var matrices []domain.Matrix
		var err error
		if len(req.CourseID) > 0 {
			matrices, err = s.FindMatrixByCourse(ctx, req.CourseID)
		} else {
			matrices, err = s.ListMatrix(ctx)
		}
		if err != nil {
			return nil, err
		}

		var list []findMatrixResponse
		for i := range matrices {
			m := matrices[i]
			list = append(list, findMatrixResponse{
				UUID:        m.UUID,
				Title:       m.Title,
				Description: m.Description,
				CreatedAt:   m.CreatedAt,
				UpdatedAt:   m.UpdatedAt,
				CourseID:    m.CourseID,
			})
		}

		return &listMatrixResponse{Matrices: list}, nil
	}
}

func decodeListMatrixRequest(_ context.Context, r *http.Request) (interface{}, error) {
	courseID := r.FormValue("course_id")
	return listMatrixRequest{CourseID: courseID}, nil
}

func encodeListMatrixResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
