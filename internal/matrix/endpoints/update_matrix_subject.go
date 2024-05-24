package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
	"github.com/sumelms/microservice-course/pkg/validator"
)

type UpdateMatrixSubjectRequest struct {
	MatrixUUID  uuid.UUID `json:"matrix_uuid"  validate:"required"`
	SubjectUUID uuid.UUID `json:"subject_uuid" validate:"required"`
	IsRequired  bool      `json:"is_required"`
	Group       string    `json:"group"        validate:"required"`
}

type UpdateMatrixSubjectResponse struct {
	MatrixSubject *MatrixSubjectResponse `json:"matrix_subject"`
}

// NewUpdateMatrixSubjectHandler updates new matrix subject handler
// @Summary      Update matrix subject
// @Description  Update a matrix subject
// @Tags         matrix_subjects
// @Accept       json
// @Produce      json
// @Param        matrix_uuid	  	path    string  true  "Matrix UUID" Format(uuid)
// @Param        subject_uuid	  	path    string  true  "Subject UUID" Format(uuid)
// @Param        matrix_subject	    body	UpdateMatrixSubjectRequest		true	"Update MatrixSubject"
// @Success      200      {object}  UpdateMatrixSubjectResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /matrices/{matrix_uuid}/subjects/{subject_uuid} [put].
func NewUpdateMatrixSubjectHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeUpdateMatrixSubjectEndpoint(s),
		decodeUpdateMatrixSubjectRequest,
		encodeUpdateMatrixSubjectResponse,
		opts...,
	)
}

//nolint:dupl
func makeUpdateMatrixSubjectEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(UpdateMatrixSubjectRequest)
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

		err := s.UpdateMatrixSubject(ctx, &matrixSubject)
		if err != nil {
			return nil, err
		}

		return &UpdateMatrixSubjectResponse{
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

func decodeUpdateMatrixSubjectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	MatrixUUID, ok := vars["matrix_uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	SubjectUUID, ok := vars["subject_uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	var req UpdateMatrixSubjectRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.MatrixUUID = uuid.MustParse(MatrixUUID)
	req.SubjectUUID = uuid.MustParse(SubjectUUID)

	return req, nil
}

func encodeUpdateMatrixSubjectResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
