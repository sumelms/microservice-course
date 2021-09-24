package endpoints

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

type listMatrixResponse struct {
	Matrices []findMatrixResponse `json:"matrices"`
}

func NewListMatrixHandler(s domain.Service, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeListMatrixEndpoint(s),
		decodeListMatrixRequest,
		encodeListMatrixResponse,
		opts...,
	)
}

func makeListMatrixEndpoint(s domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		matrices, err := s.ListMatrix(ctx)
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
			})
		}

		return &listMatrixResponse{Matrices: list}, nil
	}
}

func decodeListMatrixRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeListMatrixResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
