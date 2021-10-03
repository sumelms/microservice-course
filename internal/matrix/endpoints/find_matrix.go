package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/endpoint"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

type findMatrixRequest struct {
	UUID string `json:"uuid" validate:"required"`
}

type findMatrixResponse struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CourseID    string    `json:"course_id"`
}

func NewFindMatrixHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeFindMatrixEndpoint(s),
		decodeFindMatrixRequest,
		encodeFindMatrixResponse,
		opts...,
	)
}

func makeFindMatrixEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(findMatrixRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		m, err := s.FindMatrix(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return &findMatrixResponse{
			UUID:        m.UUID,
			Title:       m.Title,
			Description: m.Description,
			CreatedAt:   m.CreatedAt,
			UpdatedAt:   m.UpdatedAt,
			CourseID:    m.CourseID,
		}, nil
	}
}

func decodeFindMatrixRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	return findMatrixRequest{UUID: id}, nil
}

func encodeFindMatrixResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
