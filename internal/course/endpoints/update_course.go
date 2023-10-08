package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/pkg/validator"
)

type updateCourseRequest struct {
	UUID        uuid.UUID `json:"uuid"        validate:"required"`
	Code        string    `json:"code" validate:"required,max=100"`
	Name        string    `json:"name"        validate:"required,max=100"`
	Underline   string    `json:"underline"   validate:"required,max=100"`
	Image       string    `json:"image"`
	ImageCover  string    `json:"image_cover"`
	Excerpt     string    `json:"excerpt"     validate:"required,max=140"`
	Description string    `json:"description" validate:"required,max=255"`
}

type updateCourseResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Underline   string    `json:"underline"`
	Image       string    `json:"image"`
	ImageCover  string    `json:"image_cover"`
	Excerpt     string    `json:"excerpt"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

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
		req, ok := request.(updateCourseRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		c := domain.Course{}
		data, _ := json.Marshal(req)
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}

		if err := s.UpdateCourse(ctx, &c); err != nil {
			return nil, err
		}

		return updateCourseResponse{
			UUID:        c.UUID,
			Code:        c.Code,
			Name:        c.Name,
			Underline:   c.Underline,
			Image:       c.Image,
			ImageCover:  c.ImageCover,
			Excerpt:     c.Excerpt,
			Description: c.Description,
		}, nil
	}
}

func decodeUpdateCourseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	var req updateCourseRequest
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
