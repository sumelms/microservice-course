package endpoints

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-course/internal/course/domain"
)

type ListCoursesResponse struct {
	Courses []CourseResponse `json:"courses"`
}

func NewListCoursesHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeListCoursesEndpoint(s),
		decodeListCoursesRequest,
		encodeListCoursesResponse,
		opts...,
	)
}

func makeListCoursesEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		courses, err := s.Courses(ctx)
		if err != nil {
			return nil, err
		}

		var list []CourseResponse
		for i := range courses {
			c := courses[i]
			list = append(list, CourseResponse{
				UUID:        c.UUID,
				Code:        c.Code,
				Name:        c.Name,
				Underline:   c.Underline,
				Image:       c.Image,
				ImageCover:  c.ImageCover,
				Excerpt:     c.Excerpt,
				Description: c.Description,
				CreatedAt:   c.CreatedAt,
				UpdatedAt:   c.UpdatedAt,
			})
		}

		return &ListCoursesResponse{Courses: list}, nil
	}
}

func decodeListCoursesRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeListCoursesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
