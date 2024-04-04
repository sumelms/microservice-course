package endpoints

//nolint:dupl

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/course/domain"
)

type DeleteCourseRequest struct {
	UUID uuid.UUID `json:"uuid" validate:"required"`
}

type DeletedCourseResponse struct {
	UUID      uuid.UUID `json:"uuid"`
	DeletedAt time.Time `json:"deleted_at"`
}

type DeleteCourseResponse struct {
	Course *DeletedCourseResponse `json:"course"`
}

func NewDeleteCourseHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeDeleteCourseEndpoint(s),
		decodeDeleteCourseRequest,
		encodeDeleteCourseResponse,
		opts...,
	)
}

func makeDeleteCourseEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(DeleteCourseRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		deletedCourse := domain.DeletedCourse{UUID: req.UUID}
		if err := s.DeleteCourse(ctx, &deletedCourse); err != nil {
			return nil, err
		}

		return DeleteCourseResponse{
			Course: &DeletedCourseResponse{
				UUID:      deletedCourse.UUID,
				DeletedAt: deletedCourse.DeletedAt,
			},
		}, nil
	}
}

func decodeDeleteCourseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return DeleteCourseRequest{UUID: uid}, nil
}

func encodeDeleteCourseResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if response != nil {
		return kithttp.EncodeJSONResponse(ctx, w, response)
	}
	return nil
}
