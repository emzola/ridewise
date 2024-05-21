package model

import "time"

type OTPRecord struct {
	OTP    string
	Expiry time.Time
}
