package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/course/domain"
)

type getCourseRequest struct {
	UUID string `json:"uuid"`
}

type getCourseResponse struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Excerpt     string    `json:"excerpt"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewGetCourseHandler(s domain.Service, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeGetCourseEndpoint(s),
		decodeGetCourseRequest,
		encodeGetCourseResponse,
		opts...,
	)
}

func makeGetCourseEndpoint(s domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(getCourseRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		course, err := s.GetCourse(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return &getCourseResponse{
			UUID:        course.UUID,
			Title:       course.Title,
			Subtitle:    course.Subtitle,
			Excerpt:     course.Excerpt,
			Description: course.Description,
		}, nil
	}
}

func decodeGetCourseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	return getCourseRequest{UUID: id}, nil
}

func encodeGetCourseResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
