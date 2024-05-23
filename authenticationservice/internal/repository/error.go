package repository

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrInvalidOTP          = errors.New("invalid or expired OTP")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)
