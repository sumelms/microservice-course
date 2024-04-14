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

// NewListCoursesHandler list courses handler
// @Summary      List courses
// @Description  List a new courses
// @Tags         courses
// @Produce      json
// @Success      200      {object}  ListCoursesResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /courses [get].
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
			course := courses[i]
			list = append(list, CourseResponse{
				UUID:        course.UUID,
				Code:        course.Code,
				Name:        course.Name,
				Underline:   course.Underline,
				Image:       course.Image,
				ImageCover:  course.ImageCover,
				Excerpt:     course.Excerpt,
				Description: course.Description,
				CreatedAt:   course.CreatedAt,
				UpdatedAt:   course.UpdatedAt,
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
