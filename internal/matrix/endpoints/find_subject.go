package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

type FindSubjectRequest struct {
	UUID uuid.UUID `json:"uuid" validate:"required"`
}

type FindSubjectResponse struct {
	Subject *SubjectResponse `json:"subject"`
}

// NewFindSubjectHandler find subject handler
// @Summary      Find subject
// @Description  Find a new subject
// @Tags         subjects
// @Produce      json
// @Param        uuid	  path      string  true  "Subject UUID" Format(uuid)
// @Success      200      {object}  FindSubjectResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /subjects/{uuid} [get].
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
		req, ok := request.(FindSubjectRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		subject, err := s.Subject(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return &FindSubjectResponse{
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

func decodeFindSubjectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	uid := uuid.MustParse(id)

	return FindSubjectRequest{UUID: uid}, nil
}

func encodeFindSubjectResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
