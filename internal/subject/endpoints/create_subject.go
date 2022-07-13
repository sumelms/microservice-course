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

	"github.com/sumelms/microservice-course/internal/subject/domain"
	"github.com/sumelms/microservice-course/pkg/validator"
)

type createSubjectRequest struct {
	Title       string `json:"title" validate:"required,max=100"`
	Subtitle    string `json:"subtitle" validate:"required,max=100"`
	Excerpt     string `json:"excerpt" validate:"required,max=140"`
	Description string `json:"description" validate:"required,max=255"`
}

type createSubjectResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Excerpt     string    `json:"excerpt"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewCreateSubjectHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeCreateSubjectEndpoint(s),
		decodeCreateSubjectRequest,
		encodeCreateSubjectResponse,
		opts...,
	)
}

func makeCreateSubjectEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createSubjectRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		c := domain.Subject{}
		data, _ := json.Marshal(req)
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}

		if err := s.CreateSubject(ctx, &c); err != nil {
			return nil, err
		}

		return createSubjectResponse{
			UUID:      c.UUID,
			Title:     c.Title,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		}, nil
	}
}

func decodeCreateSubjectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createSubjectRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeCreateSubjectResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}