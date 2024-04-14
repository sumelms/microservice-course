package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/course/domain"
)

type FindCourseRequest struct {
	UUID uuid.UUID `json:"uuid"`
}

type FindCourseResponse struct {
	Course *CourseResponse `json:"course"`
}

// NewFindCourseHandler find course handler
// @Summary      Find course
// @Description  Find a new course
// @Tags         courses
// @Produce      json
// @Param        uuid	  path      string  true  "Course UUID" Format(uuid)
// @Success      200      {object}  FindCourseResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /courses/{uuid} [get].
func NewFindCourseHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeFindCourseEndpoint(s),
		decodeFindCourseRequest,
		encodeFindCourseResponse,
		opts...,
	)
}

func makeFindCourseEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(FindCourseRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		course, err := s.Course(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return &FindCourseResponse{
			Course: &CourseResponse{
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
			},
		}, nil
	}
}

func decodeFindCourseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	uid := uuid.MustParse(id)

	return FindCourseRequest{UUID: uid}, nil
}

func encodeFindCourseResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
