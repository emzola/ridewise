package auth

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/emzola/ridewise/authenticationservice/internal/controller"
)

type authRepository interface {
	SaveOTP(phoneNumber string, otp string, expiry time.Time) error
	GetPhoneNumberByOTP(otp string) (string, error)
	VerifyOTP(phoneNumber string, otp string) (bool, error)
	SaveRefreshToken(phoneNumber string, refreshToken string) error
}

type Controller struct {
	repo authRepository
}

func New(repo authRepository) *Controller {
	return &Controller{repo}
}

func (c *Controller) GenerateOTP(phoneNumber string) (string, error) {
	otp, err := generateSecureOTP()
	if err != nil {
		return "", err
	}
	err = c.repo.SaveOTP(phoneNumber, otp, time.Now().Add(5*time.Minute))
	if err != nil {
		return "", err
	}
	return otp, nil
}

func (c *Controller) GetPhoneNumberByOTP(otp string) (string, error) {
	phoneNumber, err := c.repo.GetPhoneNumberByOTP(otp)
	if err != nil {
		return "", controller.ErrNotFound
	}
	return phoneNumber, nil
}

func (c *Controller) VerifyOTP(phoneNumber string, otp string) (string, string, error) {
	valid, err := c.repo.VerifyOTP(phoneNumber, otp)
	if err != nil || !valid {
		return "", "", controller.ErrInvalidOTP
	}
	accessToken, err := generateToken(phoneNumber, 15*time.Minute)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := generateToken(phoneNumber, 365*7*time.Minute)
	if err != nil {
		return "", "", err
	}
	err = c.repo.SaveRefreshToken(phoneNumber, refreshToken)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
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

// generateToken generates a JWT token.
func generateToken(phoneNumber string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone": phoneNumber,
		"exp":   time.Now().Add(duration).Unix(),
	})
	return token.SignedString([]byte("secret"))
}
