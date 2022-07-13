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

	"github.com/sumelms/microservice-course/internal/subject/domain"
)

type findSubjectRequest struct {
	UUID uuid.UUID `json:"uuid"`
}

type findSubjectResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Excerpt     string    `json:"excerpt"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewFindSubjectHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeFindSubjectEndpoint(s),
		decodeFindSubjectRequest,
		encodeFindSubjectResponse,
		opts...,
	)
}

func makeFindSubjectEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(findSubjectRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		c, err := s.Subject(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return &findSubjectResponse{
			UUID:  c.UUID,
			Title: c.Title,
		}, nil
	}
}

func decodeFindSubjectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	uid := uuid.MustParse(id)

	return findSubjectRequest{UUID: uid}, nil
}

func encodeFindSubjectResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
