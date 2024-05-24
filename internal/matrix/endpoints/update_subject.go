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
	"github.com/sumelms/microservice-course/internal/matrix/domain"
	"github.com/sumelms/microservice-course/pkg/validator"
)

type UpdateSubjectRequest struct {
	UUID        uuid.UUID  `json:"uuid"         validate:"required"`
	Code        string     `json:"code"         validate:"required,max=45"`
	Name        string     `json:"name"         validate:"required,max=100"`
	Objective   string     `json:"objective"    validate:"max=245"`
	Credit      float32    `json:"credit"`
	Workload    float32    `json:"workload"`
	PublishedAt *time.Time `json:"published_at"`
}

type UpdateSubjectResponse struct {
	Subject *SubjectResponse `json:"subject"`
}

// NewUpdateSubjectHandler updates new subject handler
// @Summary      Update subject
// @Description  Update a subject
// @Tags         subjects
// @Accept       json
// @Produce      json
// @Param        uuid		path    string  				true	"Subject UUID" Format(uuid)
// @Param        subject	body	UpdateSubjectRequest	true	"Update Subject"
// @Success      200      {object}  UpdateSubjectResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /subjects/{uuid} [put].
func NewUpdateSubjectHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeUpdateSubjectEndpoint(s),
		decodeUpdateSubjectRequest,
		encodeUpdateSubjectResponse,
		opts...,
	)
}

//nolint:dupl
func makeUpdateSubjectEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(UpdateSubjectRequest)
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

		if err := s.UpdateSubject(ctx, &subject); err != nil {
			return nil, err
		}

		return UpdateSubjectResponse{
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

func decodeUpdateSubjectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	var req UpdateSubjectRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.UUID = uuid.MustParse(id)

	return req, nil
}

func encodeUpdateSubjectResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
