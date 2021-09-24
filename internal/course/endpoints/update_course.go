package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/pkg/validator"
)

type updateCourseRequest struct {
	UUID        string `json:"uuid" validate:"required"`
	Title       string `json:"title" validate:"required,max=100"`
	Subtitle    string `json:"subtitle" validade:"required,max=100"`
	Excerpt     string `json:"excerpt" validate:"required,max=140"`
	Description string `json:"description" validate:"required,max=255"`
}

type updateCourseResponse struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Excerpt     string    `json:"excerpt"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewUpdateCourseHandler(svc domain.Service, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeUpdateCourseEndpoint(svc),
		decodeUpdateCourseRequest,
		encodeUpdateCourseResponse,
		opts...,
	)
}

func makeUpdateCourseEndpoint(s domain.Service) endpoint.Endpoint {
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
		err := json.Unmarshal(data, &c)
		if err != nil {
			return nil, err
		}

		updated, err := s.UpdateCourse(ctx, &c)
		if err != nil {
			return nil, err
		}

		return updateCourseResponse{
			UUID:        updated.UUID,
			Title:       updated.Title,
			Subtitle:    updated.Subtitle,
			Excerpt:     updated.Excerpt,
			Description: updated.Description,
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

	req.UUID = id

	return req, nil
}

func encodeUpdateCourseResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}