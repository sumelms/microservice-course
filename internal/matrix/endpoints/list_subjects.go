package endpoints

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

type ListSubjectsResponse struct {
	Subjects []SubjectResponse `json:"subjects"`
}

// NewListSubjectsHandler list subjects handler
// @Summary      List subjects
// @Description  List a new subjects
// @Tags         subjects
// @Accept       json
// @Produce      json
// @Success      200      {object}  ListSubjectsResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /subjects [get].
func NewListSubjectsHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeListSubjectsEndpoint(s),
		decodeListSubjectsRequest,
		encodeListSubjectsResponse,
		opts...,
	)
}

func makeListSubjectsEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		subjects, err := s.Subjects(ctx)
		if err != nil {
			return nil, err
		}

		var list []SubjectResponse
		for i := range subjects {
			subject := subjects[i]
			list = append(list, SubjectResponse{
				UUID:        subject.UUID,
				Code:        subject.Code,
				Name:        subject.Name,
				Objective:   subject.Objective,
				Credit:      subject.Credit,
				Workload:    subject.Workload,
				CreatedAt:   subject.CreatedAt,
				UpdatedAt:   subject.UpdatedAt,
				PublishedAt: subject.PublishedAt,
			})
		}

		return &ListSubjectsResponse{Subjects: list}, nil
	}
}

func decodeListSubjectsRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeListSubjectsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
