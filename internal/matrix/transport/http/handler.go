package http

import (
	"net/http"

	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
	"github.com/sumelms/microservice-course/internal/matrix/endpoints"
	"github.com/sumelms/microservice-course/pkg/errors"
)

func NewHTTPHandler(r *mux.Router, s domain.ServiceInterface, logger log.Logger) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(errors.EncodeError),
	}

	matrixRouter := NewMatrixRouter(s, opts...)
	subjectRouter := NewSubjectRouter(s, opts...)

	r.PathPrefix("/matrices").Handler(matrixRouter)
	r.PathPrefix("/subjects").Handler(subjectRouter)
}

func NewMatrixRouter(s domain.ServiceInterface, opts ...kithttp.ServerOption) *mux.Router {
	r := mux.NewRouter().PathPrefix("/matrices").Subrouter().StrictSlash(true)

	listMatricesHandler := endpoints.NewListMatricesHandler(s, opts...)
	r.Handle("", listMatricesHandler).Methods(http.MethodGet)

	createMatrixHandler := endpoints.NewCreateMatrixHandler(s, opts...)
	r.Handle("", createMatrixHandler).Methods(http.MethodPost)

	findMatrixHandler := endpoints.NewFindMatrixHandler(s, opts...)
	r.Handle("/{uuid}", findMatrixHandler).Methods(http.MethodGet)

	updateMatrixHandler := endpoints.NewUpdateMatrixHandler(s, opts...)
	r.Handle("/{uuid}", updateMatrixHandler).Methods(http.MethodPut)

	deleteMatrixHandler := endpoints.NewDeleteMatrixHandler(s, opts...)
	r.Handle("/{uuid}", deleteMatrixHandler).Methods(http.MethodDelete)

	return r
}

func NewSubjectRouter(s domain.ServiceInterface, opts ...kithttp.ServerOption) *mux.Router {
	r := mux.NewRouter().PathPrefix("/subjects").Subrouter().StrictSlash(true)

	listSubjectHandler := endpoints.NewListSubjectHandler(s, opts...)
	r.Handle("", listSubjectHandler).Methods(http.MethodGet)

	createSubjectHandler := endpoints.NewCreateSubjectHandler(s, opts...)
	r.Handle("", createSubjectHandler).Methods(http.MethodPost)

	findSubjectHandler := endpoints.NewFindSubjectHandler(s, opts...)
	r.Handle("/{uuid}", findSubjectHandler).Methods(http.MethodGet)

	updateSubjectHandler := endpoints.NewUpdateSubjectHandler(s, opts...)
	r.Handle("/{uuid}", updateSubjectHandler).Methods(http.MethodPut)

	deleteSubjectHandler := endpoints.NewDeleteSubjectHandler(s, opts...)
	r.Handle("/{uuid}", deleteSubjectHandler).Methods(http.MethodDelete)

	return r
}
