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

type FindMatrixRequest struct {
	UUID uuid.UUID `json:"uuid" validate:"required"`
}

type FindMatrixResponse struct {
	Matrix *MatrixResponse `json:"matrix"`
}

// NewFindMatrixHandler find matrix handler
// @Summary      Find matrix
// @Description  Find a new matrix
// @Tags         matrices
// @Produce      json
// @Param        uuid	  path      string  true  "Matrix UUID" Format(uuid)
// @Success      200      {object}  FindMatrixResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /matrices/{uuid} [get].
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
		req, ok := request.(FindMatrixRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		matrix, err := s.Matrix(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return &FindMatrixResponse{
			Matrix: &MatrixResponse{
				UUID:        matrix.UUID,
				CourseUUID:  matrix.CourseUUID,
				Code:        matrix.Code,
				Name:        matrix.Name,
				Description: matrix.Description,
				CreatedAt:   matrix.CreatedAt,
				UpdatedAt:   matrix.UpdatedAt,
			},
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

	return FindMatrixRequest{UUID: uid}, nil
}

func encodeFindMatrixResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
