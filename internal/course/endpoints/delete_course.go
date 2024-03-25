package endpoints

//nolint:dupl

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/internal/shared"
)

type deleteCourseRequest struct {
	UUID uuid.UUID `json:"uuid" validate:"required"`
}

type deleteCourseResponse struct {
	Course *shared.Deleted `json:"course"`
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
		req, ok := request.(deleteCourseRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		deleted := shared.Deleted{UUID: req.UUID}
		if err := s.DeleteCourse(ctx, &deleted); err != nil {
			return nil, err
		}

		return deleteCourseResponse{
			Course: &deleted,
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

	return deleteCourseRequest{UUID: uid}, nil
}

func encodeDeleteCourseResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if response != nil {
		return kithttp.EncodeJSONResponse(ctx, w, response)
	}
	return nil
}
