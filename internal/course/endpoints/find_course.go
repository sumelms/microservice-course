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

type findCourseRequest struct {
	UUID string `json:"uuid"`
}

type findCourseResponse struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Excerpt     string    `json:"excerpt"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewFindCourseHandler(s domain.Service, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeFindCourseEndpoint(s),
		decodeFindCourseRequest,
		encodeFindCourseResponse,
		opts...,
	)
}

func makeFindCourseEndpoint(s domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(findCourseRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		c, err := s.FindCourse(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return &findCourseResponse{
			UUID:        c.UUID,
			Title:       c.Title,
			Subtitle:    c.Subtitle,
			Excerpt:     c.Excerpt,
			Description: c.Description,
		}, nil
	}
}

func decodeFindCourseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	return findCourseRequest{UUID: id}, nil
}

func encodeFindCourseResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
