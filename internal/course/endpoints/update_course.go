package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/pkg/validator"
)

type UpdateCourseRequest struct {
	UUID        uuid.UUID `json:"uuid"        validate:"required"`
	Code        string    `json:"code"        validate:"required,max=100"`
	Name        string    `json:"name"        validate:"required,max=100"`
	Underline   string    `json:"underline"   validate:"required,max=100"`
	Image       string    `json:"image"`
	ImageCover  string    `json:"image_cover"`
	Excerpt     string    `json:"excerpt"     validate:"required,max=140"`
	Description string    `json:"description" validate:"required,max=255"`
}

type UpdateCourseResponse struct {
	Course *CourseResponse `json:"course"`
}

// NewUpdateCourseHandler updates new course handler
// @Summary      Update course
// @Description  Update a course
// @Tags         courses
// @Accept       json
// @Produce      json
// @Param        uuid	  path      string  true  "Course UUID" Format(uuid)
// @Param        course	  body		UpdateCourseRequest		true	"Update Course"
// @Success      200      {object}  UpdateCourseResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /courses/{uuid} [put].
func NewUpdateCourseHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeUpdateCourseEndpoint(s),
		decodeUpdateCourseRequest,
		encodeUpdateCourseResponse,
		opts...,
	)
}

func makeUpdateCourseEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(UpdateCourseRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		course := domain.Course{}
		data, _ := json.Marshal(req)
		if err := json.Unmarshal(data, &course); err != nil {
			return nil, err
		}

		if err := s.UpdateCourse(ctx, &course); err != nil {
			return nil, err
		}

		return UpdateCourseResponse{
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

func decodeUpdateCourseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	var req UpdateCourseRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.UUID = uuid.MustParse(id)

	return req, nil
}

func encodeUpdateCourseResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
