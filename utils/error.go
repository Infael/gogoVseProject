package utils

import (
	"fmt"
	"net/http"
	"runtime"
)

type StatusError struct {
	Code   int
	Err    error
	Caller string
}

// implement the error interface
func (e StatusError) Error() string {
	return e.Err.Error()
}

func (e StatusError) StatusCode() int {
	return e.Code
}

func NewError(err error, code int) StatusError {
	pc, _, line, _ := runtime.Caller(2)
	caller := runtime.FuncForPC(pc).Name()

	return StatusError{Err: err, Code: code, Caller: fmt.Sprintf("%s:%d", caller, line)}
}

func ErrorBadRequest(err error) StatusError {
	return NewError(err, http.StatusBadRequest)
}

func ErrorUnauthorized(err error) StatusError {
	return NewError(err, http.StatusUnauthorized)
}

func ErrorForbidden(err error) StatusError {
	return NewError(err, http.StatusForbidden)
}

func ErrorNotFound(err error) StatusError {
	return NewError(err, http.StatusNotFound)
}

func InternalServerError(err error) StatusError {
	return NewError(err, http.StatusInternalServerError)
}
