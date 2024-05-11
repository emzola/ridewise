package repository

import "errors"

var (
	ErrNotFound       = errors.New("the requested resource could not be found")
	ErrDuplicatePhone = errors.New("a user with this phone number already exists")
	ErrDuplicateEmail = errors.New("a user with the email address already exists")
)
