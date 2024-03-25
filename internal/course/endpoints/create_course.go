package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/pkg/validator"
)

type createCourseRequest struct {
	Code        string `json:"code"        validate:"required,max=15"`
	Name        string `json:"name"        validate:"required,max=100"`
	Underline   string `json:"underline"   validate:"required,max=100"`
	Image       string `json:"image"`
	ImageCover  string `json:"image_cover"`
	Excerpt     string `json:"excerpt"     validate:"required,max=140"`
	Description string `json:"description" validate:"required,max=255"`
}

type createCourseResponse struct {
	Course *domain.Course `json:"course"`
}

// NewCreateCourseHandler creates new course handler
// @Summary      Create course
// @Description  Create a new course
// @Tags         course
// @Accept       json
// @Produce      json
// @Param        course	  body		createCourseRequest		true	"Add Course"
// @Success      200      {object}  createCourseResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /courses [post].
func NewCreateCourseHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeCreateCourseEndpoint(s),
		decodeCreateCourseRequest,
		encodeCreateCourseResponse,
		opts...,
	)
}

func makeCreateCourseEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createCourseRequest)
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

		if err := s.CreateCourse(ctx, &course); err != nil {
			return nil, err
		}

		return createCourseResponse{
			Course: &course,
		}, nil
	}
}

func decodeCreateCourseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createCourseRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func encodeCreateCourseResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
