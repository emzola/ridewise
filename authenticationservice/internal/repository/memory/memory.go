package memory

import (
	"sync"
	"time"

	"github.com/emzola/ridewise/authenticationservice/pkg/model"
)

type Repository struct {
	sync.RWMutex
	otps          map[string]model.OTPRecord
	refreshTokens map[string]string
}

func New() *Repository {
	return &Repository{
		otps:          map[string]model.OTPRecord{},
		refreshTokens: map[string]string{},
	}
}

func (r *Repository) GenerateOTP(phone string, otp string, expiry time.Time) error {
	r.Lock()
	defer r.Unlock()
	r.otps[phone] = model.OTPRecord{OTP: otp, Expiry: expiry}
	return nil
}
