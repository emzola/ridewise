package memory

import (
	"sync"
	"time"

	"github.com/emzola/ridewise/authenticationservice/internal/repository"
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

func (r *Repository) SaveOTP(phone string, otp string, expiry time.Time) error {
	r.Lock()
	defer r.Unlock()
	r.otps[phone] = model.OTPRecord{OTP: otp, Expiry: expiry}
	return nil
}

func (r *Repository) GetPhoneNumberByOTP(otp string) (string, error) {
	r.RLock()
	defer r.RUnlock()
	var phoneNumber string
	for phone, otpRecord := range r.otps {
		if otpRecord.OTP == otp {
			phoneNumber = phone
		}
	}
	if phoneNumber == "" {
		return "", repository.ErrNotFound
	}
	return phoneNumber, nil
}

func (r *Repository) VerifyOTP(phone string, otp string) (bool, error) {
	r.RLock()
	defer r.RUnlock()
	record, ok := r.otps[phone]
	if !ok || record.OTP != otp || time.Now().After(record.Expiry) {
		return false, repository.ErrInvalidOTP
	}
	return true, nil
}

func (r *Repository) SaveRefreshToken(phone string, refreshToken string) error {
	r.Lock()
	defer r.Unlock()
	r.refreshTokens[refreshToken] = phone
	return nil
}
