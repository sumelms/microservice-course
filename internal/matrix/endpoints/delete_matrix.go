package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

type DeleteMatrixRequest struct {
	UUID uuid.UUID `json:"uuid" validate:"required"`
}

type DeletedMatrixResponse struct {
	UUID      uuid.UUID `json:"uuid"`
	DeletedAt time.Time `json:"deleted_at"`
}

type DeleteMatrixResponse struct {
	Matrix *DeletedMatrixResponse `json:"matrix"`
}

func NewDeleteMatrixHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeDeleteMatrixEndpoint(s),
		decodeDeleteMatrixRequest,
		encodeDeleteMatrixResponse,
		opts...,
	)
}

func makeDeleteMatrixEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(DeleteMatrixRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		deletedMatrix := domain.DeletedMatrix{UUID: req.UUID}
		if err := s.DeleteMatrix(ctx, &deletedMatrix); err != nil {
			return nil, err
		}

		return DeleteMatrixResponse{
			Matrix: &DeletedMatrixResponse{
				UUID:      deletedMatrix.UUID,
				DeletedAt: deletedMatrix.DeletedAt,
			},
		}, nil
	}
}

func decodeDeleteMatrixRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	uid := uuid.MustParse(id)

	return DeleteMatrixRequest{UUID: uid}, nil
}

func encodeDeleteMatrixResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
