package errors

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// EncodeError encodes errors from business-logic
func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	resp := &ErrorResponse{Error: err.Error()}
	code := http.StatusInternalServerError

	var ierr *Error
	if !errors.As(err, &ierr) {
		resp.Error = "internal error"
	} else {
		switch ierr.Code() {
		case ErrCodeNotFound:
			code = http.StatusNotFound
		case ErrCodeInvalidArgument:
			code = http.StatusBadRequest
		case ErrCodeUnknown:
			code = http.StatusInternalServerError
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	content, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)

	if _, err := w.Write(content); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
