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

type ListMatrixSubjectsRequest struct {
	MatrixUUID uuid.UUID `json:"matrix_uuid" validate:"required"`
}

type ListMatrixSubjectsResponse struct {
	MatrixSubjects []MatrixSubjectResponse `json:"matrix_subjects"`
}

// NewListMatrixSubjectsHandler list matrix subjects handler
// @Summary      List matrix subjects
// @Description  List matrix subjects
// @Tags         matrix_subjects
// @Accept       json
// @Produce      json
// @Param        matrix_uuid	  path      string  true  "Matrix UUID" Format(uuid)
// @Success      200      {object}  ListMatrixSubjectsResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /matrices/{matrix_uuid}/subjects [get].
func NewListMatrixSubjectsHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeListMatrixSubjectsEndpoint(s),
		decodeListMatrixSubjectsRequest,
		encodeListMatrixSubjectsResponse,
		opts...,
	)
}

func makeListMatrixSubjectsEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(ListMatrixSubjectsRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		matrixSubjects, err := s.MatrixSubjects(ctx, req.MatrixUUID)
		if err != nil {
			return nil, err
		}

		var list []MatrixSubjectResponse
		for i := range matrixSubjects {
			matrix := matrixSubjects[i]
			list = append(list, MatrixSubjectResponse{
				UUID:        matrix.UUID,
				MatrixUUID:  matrix.MatrixUUID,
				SubjectUUID: matrix.SubjectUUID,
				IsRequired:  matrix.IsRequired,
				Group:       matrix.Group,
				CreatedAt:   matrix.CreatedAt,
				UpdatedAt:   matrix.UpdatedAt,
			})
		}

		return &ListMatrixSubjectsResponse{MatrixSubjects: list}, nil
	}
}

func decodeListMatrixSubjectsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	MatrixUUID, ok := vars["matrix_uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	uid := uuid.MustParse(MatrixUUID)

	return ListMatrixSubjectsRequest{MatrixUUID: uid}, nil
}

func encodeListMatrixSubjectsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
