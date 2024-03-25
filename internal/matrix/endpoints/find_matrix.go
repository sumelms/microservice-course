package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

type findMatrixRequest struct {
	UUID uuid.UUID `json:"uuid" validate:"required"`
}

type findMatrixResponse struct {
	Matrix *domain.Matrix `json:"matrix"`
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

		matrix, err := s.Matrix(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return &findMatrixResponse{
			Matrix: &matrix,
		}, nil
	}
}

func decodeFindMatrixRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	uid := uuid.MustParse(id)

	return findMatrixRequest{UUID: uid}, nil
}

func encodeFindMatrixResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
