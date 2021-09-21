package errors

import "fmt"

type Error struct {
	original error
	msg      string
	code     ErrorCode
}

type ErrorCode int

const (
	ErrCodeUnknown ErrorCode = iota
	ErrCodeNotFound
	ErrCodeInvalidArgument
)

func WrapErrorf(original error, code ErrorCode, format string, a ...interface{}) error {
	return &Error{
		original: original,
		code:     code,
		msg:      fmt.Sprintf(format, a...),
	}
}

func NewErrorf(code ErrorCode, format string, a ...interface{}) error {
	return WrapErrorf(nil, code, format, a...)
}

func (e *Error) Error() string {
	if e.original != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.original)
	}

	return e.msg
}

func (e *Error) Unwrap() error {
	return e.original
}

func (e *Error) Code() ErrorCode {
	return e.code
}
