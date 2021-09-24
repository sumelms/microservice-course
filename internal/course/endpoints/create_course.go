package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/pkg/validator"
)

type createCourseRequest struct {
	Title       string `json:"title" validate:"required,max=100"`
	Subtitle    string `json:"subtitle" validate:"required,max=100"`
	Excerpt     string `json:"excerpt" validate:"required,max=140"`
	Description string `json:"description" validate:"required,max=255"`
}

type createCourseResponse struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Excerpt     string    `json:"excerpt"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewCreateCourseHandler(s domain.Service, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeCreateCourseEndpoint(s),
		decodeCreateCourseRequest,
		encodeCreateCourseResponse,
		opts...,
	)
}

func makeCreateCourseEndpoint(s domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createCourseRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		c := domain.Course{}
		data, _ := json.Marshal(req)
		err := json.Unmarshal(data, &c)
		if err != nil {
			return nil, err
		}

		created, err := s.CreateCourse(ctx, &c)
		if err != nil {
			return nil, err
		}

		return createCourseResponse{
			UUID:        created.UUID,
			Title:       created.Title,
			Subtitle:    created.Subtitle,
			Excerpt:     created.Excerpt,
			Description: created.Description,
			CreatedAt:   created.CreatedAt,
			UpdatedAt:   created.UpdatedAt,
		}, err
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
