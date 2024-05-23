package memory

import (
	"context"
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

func (r *Repository) SaveOTP(ctx context.Context, phoneNumber string, otp string, expiry time.Time) error {
	r.Lock()
	defer r.Unlock()
	r.otps[phoneNumber] = model.OTPRecord{OTP: otp, Expiry: expiry}
	return nil
}

func (r *Repository) GetPhoneNumberByOTP(ctx context.Context, otp string) (string, error) {
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

func (r *Repository) VerifyOTP(ctx context.Context, phone string, otp string) (bool, error) {
	r.RLock()
	defer r.RUnlock()
	record, ok := r.otps[phone]
	if !ok || record.OTP != otp || time.Now().After(record.Expiry) {
		return false, repository.ErrInvalidOTP
	}
	return true, nil
}

func (r *Repository) SaveRefreshToken(ctx context.Context, phoneNumber string, refreshToken string) error {
	r.Lock()
	defer r.Unlock()
	r.refreshTokens[refreshToken] = phoneNumber
	return nil
}

func (r *Repository) GetPhoneNumberByRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	r.RLock()
	defer r.RUnlock()
	phoneNumber, ok := r.refreshTokens[refreshToken]
	if !ok {
		return "", repository.ErrNotFound
	}
	return phoneNumber, nil
}

func (r *Repository) DeleteRefreshToken(ctx context.Context, refreshToken string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.refreshTokens[refreshToken]; !ok {
		return repository.ErrNotFound
	}
	delete(r.refreshTokens, refreshToken)
	return nil
}
