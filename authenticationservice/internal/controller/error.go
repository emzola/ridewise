package controller

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrInvalidRequest      = errors.New("the request is invalid")
	ErrInvalidOTP          = errors.New("invalid or expired OTP")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)
