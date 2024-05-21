package auth

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"
)

type authRepository interface {
	GenerateOTP(phone string, otp string, expiry time.Time) error
}

type Controller struct {
	repo authRepository
}

func New(repo authRepository) *Controller {
	return &Controller{repo}
}

func (c *Controller) GenerateOTP(phone string) (string, error) {
	otp, err := generateSecureOTP()
	if err != nil {
		return "", err
	}
	err = c.repo.GenerateOTP(phone, otp, time.Now().Add(5*time.Minute))
	if err != nil {
		return "", err
	}
	return otp, nil
}

// generateSecureOTP generates a secure 6-digit OTP.
func generateSecureOTP() (string, error) {
	var n uint32
	err := binary.Read(rand.Reader, binary.BigEndian, &n)
	if err != nil {
		return "", fmt.Errorf("failed to generate OTP: %w", err)
	}
	otp := n % 1000000 // Ensure it's a 6-digit number
	return fmt.Sprintf("%06d", otp), nil
}
