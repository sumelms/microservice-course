package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/course/domain"
)

type findCourseRequest struct {
	UUID uuid.UUID `json:"uuid"`
}

type findCourseResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Underline   string    `json:"underline"`
	Image       string    `json:"image,omitempty"`
	ImageCover  string    `json:"image_cover,omitempty"`
	Excerpt     string    `json:"excerpt"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

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
		req, ok := request.(findCourseRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		c, err := s.Course(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return &findCourseResponse{
			UUID:        c.UUID,
			Name:        c.Name,
			Underline:   c.Underline,
			Image:       c.Image,
			ImageCover:  c.ImageCover,
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

	uid := uuid.MustParse(id)

	return findCourseRequest{UUID: uid}, nil
}

func encodeFindCourseResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
