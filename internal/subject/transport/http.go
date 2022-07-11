package transport

import (
	"github.com/go-kit/log"
	"github.com/gorilla/mux"

	"github.com/sumelms/microservice-course/internal/subject/domain"
)

func NewHTTPHandler(r *mux.Router, s domain.ServiceInterface, logger log.Logger) {
	panic("implement me")
}
