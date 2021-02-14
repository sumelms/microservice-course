package http

import (
	"context"
	"net/http"

	router "github.com/sumelms/microkit/http"
)

// NewHTTPServer creates http server router
func NewHTTPServer(ctx context.Context) http.Handler {
	r := router.NewRouter()
	return r
}
