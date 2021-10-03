package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

type findMatrixByCourseRequest struct {
	CourseID string `json:"course_id" validate:"required"`
}

type findMatrixByCourseResponse struct {
	Matrices []findMatrixResponse `json:"matrices"`
}

func NewFindMatrixByCourse(s domain.Service, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeFindMatrixByCourseEndpoint(s),
		decodeFindMatrixByCourseRequest,
		encodeFindMatrixByCourseResponse,
		opts...,
	)
}

func makeFindMatrixByCourseEndpoint(s domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(findMatrixByCourseRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		matrices, err := s.FindMatrixByCourse(ctx, req.CourseID)
		if err != nil {
			return nil, err
		}

		var list []findMatrixResponse
		for _, m := range matrices {
			list = append(list, findMatrixResponse{
				UUID:        m.UUID,
				Title:       m.Title,
				Description: m.Description,
				CreatedAt:   m.CreatedAt,
				UpdatedAt:   m.UpdatedAt,
				CourseID:    m.CourseID,
			})
		}

		return &findMatrixByCourseResponse{Matrices: list}, nil
	}
}

func decodeFindMatrixByCourseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	return findMatrixByCourseRequest{CourseID: id}, nil
}

func encodeFindMatrixByCourseResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
