package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

type ListMatricesRequest struct {
	CourseUUID uuid.UUID `json:"course_uuid,omitempty"`
}

type ListMatricesResponse struct {
	Matrices []MatrixResponse `json:"matrices"`
}

// NewListMatricesHandler list matrices handler
// @Summary      List matrices
// @Description  List a new matrices
// @Tags         matrices
// @Accept       json
// @Produce      json
// @Param        course_uuid    query     string  false  "course search by uuid"  Format(uuid)
// @Success      200      {object}  ListMatricesResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /matrices [get].
func NewListMatricesHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeListMatricesEndpoint(s),
		decodeListMatricesRequest,
		encodeListMatricesResponse,
		opts...,
	)
}

func makeListMatricesEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(ListMatricesRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		filters := &domain.MatrixFilters{}
		if req.CourseUUID != uuid.Nil {
			filters.CourseUUID = req.CourseUUID
		}

		matrices, err := s.Matrices(ctx, filters)
		if err != nil {
			return nil, err
		}

		var list []MatrixResponse
		for i := range matrices {
			matrix := matrices[i]
			list = append(list, MatrixResponse{
				UUID:        matrix.UUID,
				CourseUUID:  matrix.CourseUUID,
				Code:        matrix.Code,
				Name:        matrix.Name,
				Description: matrix.Description,
				CreatedAt:   matrix.CreatedAt,
				UpdatedAt:   matrix.UpdatedAt,
			})
		}

		return &ListMatricesResponse{Matrices: list}, nil
	}
}

func decodeListMatricesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	courseUUID := r.FormValue("course_uuid")

	request := ListMatricesRequest{}
	if len(courseUUID) > 0 {
		request.CourseUUID = uuid.MustParse(courseUUID)
	}
	return request, nil
}

func encodeListMatricesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
