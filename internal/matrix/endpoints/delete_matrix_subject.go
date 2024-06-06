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
)

type DeleteMatrixSubjectRequest struct {
	MatrixUUID  uuid.UUID `json:"matrix_uuid"  validate:"required"`
	SubjectUUID uuid.UUID `json:"subject_uuid" validate:"required"`
}

type DeletedMatrixSubjectResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	MatrixUUID  uuid.UUID `json:"matrix_uuid"`
	SubjectUUID uuid.UUID `json:"subject_uuid"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type DeleteMatrixSubjectResponse struct {
	MatrixSubject *DeletedMatrixSubjectResponse `json:"matrix_subject"`
}

// NewDeleteMatrixSubjectHandler deletes matrix subject handler
// @Summary      Delete matrix subject
// @Description  Delete a new matrix subject
// @Tags         matrix_subjects
// @Produce      json
// @Param        matrix_uuid	  path      string  true  "Matrix UUID" Format(uuid)
// @Param        subject_uuid	  path      string  true  "Subject UUID" Format(uuid)
// @Success      200      {object}  DeletedMatrixSubjectResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /matrices/{matrix_uuid}/subjects/{subject_uuid} [delete].
func NewDeleteMatrixSubjectHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeDeleteMatrixSubjectEndpoint(s),
		decodeDeleteMatrixSubjectRequest,
		encodeDeleteMatrixSubjectResponse,
		opts...,
	)
}

func makeDeleteMatrixSubjectEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(DeleteMatrixSubjectRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		var deletedMatrixSubject domain.DeletedMatrixSubject
		data, _ := json.Marshal(req)
		if err := json.Unmarshal(data, &deletedMatrixSubject); err != nil {
			return nil, err
		}

		if err := s.DeleteMatrixSubject(ctx, &deletedMatrixSubject); err != nil {
			return nil, err
		}

		return &DeleteMatrixSubjectResponse{
			MatrixSubject: &DeletedMatrixSubjectResponse{
				UUID:        deletedMatrixSubject.UUID,
				SubjectUUID: deletedMatrixSubject.SubjectUUID,
				MatrixUUID:  deletedMatrixSubject.MatrixUUID,
				DeletedAt:   deletedMatrixSubject.DeletedAt,
			},
		}, nil
	}
}

func decodeDeleteMatrixSubjectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	MatrixUUID, ok := vars["matrix_uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument: missing matrix_uuid")
	}

	SubjectUUID, ok := vars["subject_uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument: missing subject_uuid")
	}

	var req DeleteMatrixSubjectRequest

	req.MatrixUUID = uuid.MustParse(MatrixUUID)
	req.SubjectUUID = uuid.MustParse(SubjectUUID)

	return req, nil
}

func encodeDeleteMatrixSubjectResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
