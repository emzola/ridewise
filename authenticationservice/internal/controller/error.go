package controller

import "errors"

var (
	ErrNotFound       = errors.New("the requested resource could not be found")
	ErrInvalidRequest = errors.New("the request is invalid")
)
