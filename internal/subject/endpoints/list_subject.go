package endpoints

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/sumelms/microservice-course/internal/subject/domain"
)

type listSubjectResponse struct {
	Subjects []findSubjectResponse `json:"subjects"`
}

func NewListSubjectHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeListSubjectEndpoint(s),
		decodeListSubjectRequest,
		encodeListSubjectResponse,
		opts...,
	)
}

func makeListSubjectEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		cc, err := s.Subjects(ctx)
		if err != nil {
			return nil, err
		}

		var list []findSubjectResponse
		for i := range cc {
			c := cc[i]
			list = append(list, findSubjectResponse{
				UUID:      c.UUID,
				Title:     c.Title,
				CreatedAt: c.CreatedAt,
				UpdatedAt: c.UpdatedAt,
			})
		}

		return &listSubjectResponse{Subjects: list}, nil
	}
}

func decodeListSubjectRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeListSubjectResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
