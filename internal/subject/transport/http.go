package transport

import (
	"net/http"

	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"

	"github.com/sumelms/microservice-course/internal/subject/domain"
	"github.com/sumelms/microservice-course/internal/subject/endpoints"
	"github.com/sumelms/microservice-course/pkg/errors"
)

func NewHTTPHandler(r *mux.Router, s domain.ServiceInterface, logger log.Logger) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(errors.EncodeError),
	}

	listSubjectHandler := endpoints.NewListSubjectHandler(s, opts...)
	createSubjectHandler := endpoints.NewCreateSubjectHandler(s, opts...)
	findSubjectHandler := endpoints.NewFindSubjectHandler(s, opts...)
	updateSubjectHandler := endpoints.NewUpdateSubjectHandler(s, opts...)
	deleteSubjectHandler := endpoints.NewDeleteSubjectHandler(s, opts...)

	r.Handle("/subjects", createSubjectHandler).Methods(http.MethodPost)
	r.Handle("/subjects", listSubjectHandler).Methods(http.MethodGet)
	r.Handle("/subjects/{uuid}", findSubjectHandler).Methods(http.MethodGet)
	r.Handle("/subjects/{uuid}", updateSubjectHandler).Methods(http.MethodPut)
	r.Handle("/subjects/{uuid}", deleteSubjectHandler).Methods(http.MethodDelete)
}
