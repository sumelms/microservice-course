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
	"github.com/sumelms/microservice-course/internal/matrix/domain"
	"github.com/sumelms/microservice-course/pkg/validator"
)

type CreateSubjectRequest struct {
	Code        string     `json:"code"         validate:"required,max=45"`
	Name        string     `json:"name"         validate:"required,max=100"`
	Objective   string     `json:"objective"    validate:"max=245"`
	Credit      float32    `json:"credit"`
	Workload    float32    `json:"workload"`
	PublishedAt *time.Time `json:"published_at"`
}

type SubjectResponse struct {
	UUID        uuid.UUID  `json:"uuid"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Objective   string     `json:"objective,omitempty"`
	Credit      float32    `json:"credit,omitempty"`
	Workload    float32    `json:"workload,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
}

type CreateSubjectResponse struct {
	Subject *SubjectResponse `json:"subject"`
}

// NewCreateSubjectHandler creates new subject handler
// @Summary      Create subject
// @Description  Create a new subject
// @Tags         subjects
// @Accept       json
// @Produce      json
// @Param        subjects  body		CreateSubjectRequest	true	"Add Subject"
// @Success      200      {object}  CreateSubjectResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /subjects [post].
func NewCreateSubjectHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeCreateSubjectEndpoint(s),
		decodeCreateSubjectRequest,
		encodeCreateSubjectResponse,
		opts...,
	)
}

//nolint:dupl
func makeCreateSubjectEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(CreateSubjectRequest)

		fmt.Println(req, ok)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		subject := domain.Subject{}
		data, _ := json.Marshal(req)
		if err := json.Unmarshal(data, &subject); err != nil {
			return nil, err
		}

		if err := s.CreateSubject(ctx, &subject); err != nil {
			return nil, err
		}

		return CreateSubjectResponse{
			Subject: &SubjectResponse{
				UUID:        subject.UUID,
				Code:        subject.Code,
				Name:        subject.Name,
				Objective:   subject.Objective,
				Credit:      subject.Credit,
				Workload:    subject.Workload,
				CreatedAt:   subject.CreatedAt,
				UpdatedAt:   subject.UpdatedAt,
				PublishedAt: subject.PublishedAt,
			},
		}, nil
	}
}

func decodeCreateSubjectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateSubjectRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeCreateSubjectResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
