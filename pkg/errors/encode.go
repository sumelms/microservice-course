package errors

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Message any `json:"message"`
	Error   any `json:"error"`
}

type ValidatorError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// EncodeError encodes errors from business-logic.
func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	resp := &ErrorResponse{
		Message: "Invalid error",
		Error:   err.Error(),
	}

	code := http.StatusInternalServerError

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		msg := make([]ValidatorError, len(ve))
		for i, fe := range ve {
			msg[i] = ValidatorError{
				Field:   fe.Field(),
				Message: msgForTag(fe.Tag()),
			}
		}
		resp.Message = "invalid input"
		resp.Message = msg
		code = http.StatusBadRequest
	}

	var ierr *Error
	if errors.As(err, &ierr) {
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

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}

	return ""
}
