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

type DeleteSubjectRequest struct {
	UUID uuid.UUID `json:"uuid" validate:"required"`
}

type DeletedSubjectResponse struct {
	UUID      uuid.UUID `json:"uuid"`
	DeletedAt time.Time `json:"deleted_at"`
}

type DeleteSubjectResponse struct {
	Subject *DeletedSubjectResponse `json:"subject"`
}

// NewDeleteSubjectHandler deletes subject handler
// @Summary      Delete subject
// @Description  Delete a new subject
// @Tags         subjects
// @Produce      json
// @Param        uuid	  path      string  true  "Subject UUID" Format(uuid)
// @Success      200      {object}  DeleteSubjectResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /subjects/{uuid} [delete].
func NewDeleteSubjectHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeDeleteSubjectEndpoint(s),
		decodeDeleteSubjectRequest,
		encodeDeleteSubjectResponse,
		opts...,
	)
}

func makeDeleteSubjectEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(DeleteSubjectRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		deletedSubject := domain.DeletedSubject{}
		data, _ := json.Marshal(req)
		if err := json.Unmarshal(data, &deletedSubject); err != nil {
			return nil, err
		}

		if err := s.DeleteSubject(ctx, &deletedSubject); err != nil {
			return nil, err
		}

		return DeleteSubjectResponse{
			Subject: &DeletedSubjectResponse{
				UUID:      deletedSubject.UUID,
				DeletedAt: deletedSubject.DeletedAt,
			},
		}, nil
	}
}

func decodeDeleteSubjectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	uid := uuid.MustParse(id)

	return DeleteSubjectRequest{UUID: uid}, nil
}

func encodeDeleteSubjectResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
