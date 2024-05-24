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

type CreateMatrixSubjectRequest struct {
	MatrixUUID  uuid.UUID `json:"matrix_uuid"  validate:"required"`
	SubjectUUID uuid.UUID `json:"subject_uuid" validate:"required"`
	IsRequired  bool      `json:"is_required"`
	Group       string    `json:"group"        validate:"required"`
}

type MatrixSubjectResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	MatrixUUID  uuid.UUID `json:"matrix_uuid"`
	SubjectUUID uuid.UUID `json:"subject_uuid"`
	IsRequired  bool      `json:"is_required"`
	Group       string    `json:"group"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateMatrixSubjectResponse struct {
	MatrixSubject *MatrixSubjectResponse `json:"matrix_subject"`
}

// NewCreateMatrixSubjectHandler creates matrix subject handler
// @Summary      Create matrix subject
// @Description  Create a new matrix subject
// @Tags         matrix_subjects
// @Produce      json
// @Param        matrix_uuid	  path      string  					true  "Matrix UUID" Format(uuid)
// @Param        matrix_subject	  body		CreateMatrixSubjectRequest	true  "Add Matrix Subject"
// @Success      200      {object}  CreateMatrixSubjectResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /matrices/{matrix_uuid}/subjects/ [post].
func NewCreateMatrixSubjectHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeCreateMatrixSubjectEndpoint(s),
		decodeCreateMatrixSubjectRequest,
		encodeCreateMatrixSubjectResponse,
		opts...,
	)
}

func makeCreateMatrixSubjectEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(CreateMatrixSubjectRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		var matrixSubject domain.MatrixSubject
		data, _ := json.Marshal(req)
		if err := json.Unmarshal(data, &matrixSubject); err != nil {
			return nil, err
		}

		if err := s.CreateMatrixSubject(ctx, &matrixSubject); err != nil {
			return nil, err
		}

		return &CreateMatrixSubjectResponse{
			MatrixSubject: &MatrixSubjectResponse{
				UUID:        matrixSubject.UUID,
				MatrixUUID:  matrixSubject.MatrixUUID,
				SubjectUUID: matrixSubject.SubjectUUID,
				IsRequired:  matrixSubject.IsRequired,
				Group:       matrixSubject.Group,
				CreatedAt:   matrixSubject.CreatedAt,
				UpdatedAt:   matrixSubject.UpdatedAt,
			},
		}, nil
	}
}

func decodeCreateMatrixSubjectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	MatrixUUID, ok := vars["matrix_uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	var req CreateMatrixSubjectRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.MatrixUUID = uuid.MustParse(MatrixUUID)

	return req, nil
}

func encodeCreateMatrixSubjectResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
