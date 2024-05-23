package repository

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrDuplicatePhone = errors.New("phone number already exists")
	ErrDuplicateEmail = errors.New("email already exists")
)
