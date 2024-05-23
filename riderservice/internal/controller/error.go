package controller

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrInvalidRequest = errors.New("the request is invalid")
	ErrDuplicatePhone = errors.New("phone number already exists")
	ErrDuplicateEmail = errors.New("email already exists")
)
