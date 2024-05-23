package repository

import "errors"

var (
	ErrNotFound   = errors.New("the requested resource could not be found")
	ErrInvalidOTP = errors.New("invalid or expired OTP")
)
