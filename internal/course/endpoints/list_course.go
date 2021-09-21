package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-course/internal/course/domain"
	"net/http"
)

type listCourseResponse struct {
	Courses []getCourseResponse `json:"courses"`
}

func NewListCourseHandler(s domain.Service, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeListCourseEndpoint(s),
		decodeListCourseRequest,
		encodeListCourseResponse,
		opts...,
	)
}

func makeListCourseEndpoint(s domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		courses, err := s.ListCourse(ctx)
		if err != nil {
			return nil, err
		}

		cr := make([]getCourseResponse, len(courses))
		for i := range courses {
			c := courses[i]
			cr[i] = getCourseResponse{
				UUID:        c.UUID,
				Title:       c.Title,
				Subtitle:    c.Subtitle,
				Excerpt:     c.Excerpt,
				Description: c.Description,
				CreatedAt:   c.CreatedAt,
				UpdatedAt:   c.UpdatedAt,
			}
		}

		return &listCourseResponse{
			Courses: cr,
		}, nil
	}
}

func decodeListCourseRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeListCourseResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
