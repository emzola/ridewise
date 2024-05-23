package auth

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/emzola/ridewise/authenticationservice/internal/controller"
	"github.com/emzola/ridewise/authenticationservice/internal/repository"
)

const (
	accessTokenDuration  = 15 * time.Minute
	refreshTokenDuration = 365 * 7 * time.Minute
)

type authRepository interface {
	SaveOTP(ctx context.Context, phoneNumber string, otp string, expiry time.Time) error
	GetPhoneNumberByOTP(ctx context.Context, otp string) (string, error)
	VerifyOTP(ctx context.Context, phoneNumber string, otp string) (bool, error)
	SaveRefreshToken(ctx context.Context, phoneNumber string, refreshToken string) error
	GetPhoneNumberByRefreshToken(ctx context.Context, refreshToken string) (string, error)
	DeleteRefreshToken(ctx context.Context, refreshToken string) error
}

type Controller struct {
	repo authRepository
}

func New(repo authRepository) *Controller {
	return &Controller{repo}
}

func (c *Controller) GenerateOTP(ctx context.Context, phoneNumber string) (string, error) {
	otp, err := generateSecureOTP()
	if err != nil {
		return "", err
	}
	err = c.repo.SaveOTP(ctx, phoneNumber, otp, time.Now().Add(5*time.Minute))
	if err != nil {
		return "", err
	}
	return otp, nil
}

func (c *Controller) GetPhoneNumberByOTP(ctx context.Context, otp string) (string, error) {
	phoneNumber, err := c.repo.GetPhoneNumberByOTP(ctx, otp)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return "", controller.ErrNotFound
		default:
			return "", err
		}
	}
	return phoneNumber, nil
}

func (c *Controller) VerifyOTP(ctx context.Context, phoneNumber string, otp string) (string, string, error) {
	valid, err := c.repo.VerifyOTP(ctx, phoneNumber, otp)
	if err != nil || !valid {
		return "", "", controller.ErrInvalidOTP
	}
	accessToken, err := generateToken(phoneNumber, accessTokenDuration)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := generateToken(phoneNumber, refreshTokenDuration)
	if err != nil {
		return "", "", err
	}
	err = c.repo.SaveRefreshToken(ctx, phoneNumber, refreshToken)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (c *Controller) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	phoneNumber, err := c.repo.GetPhoneNumberByRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", controller.ErrInvalidRefreshToken
	}
	return generateToken(phoneNumber, accessTokenDuration)
}

func (c *Controller) DeleteRefreshToken(ctx context.Context, refreshToken string) error {
	err := c.repo.DeleteRefreshToken(ctx, refreshToken)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrNotFound):
			return controller.ErrNotFound
		default:
			return err
		}
	}
	return nil
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
	return token.SignedString([]byte(os.Getenv("SECRET")))
}
