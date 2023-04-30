package errors

import "errors"

var ErrBadRequest = errors.New("the provided input is invalid")
var ErrInternalServerError = errors.New("internal server error")
var ErrNotFound = errors.New("not found")
