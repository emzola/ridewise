package gateway

import (
	"context"
	"fmt"
	"os"
)

type authenticationServiceGateway interface {
	GenerateOTP(ctx context.Context, phoneNumber string) (string, error)
	// GetPhoneNumberByOTP(ctx context.Context, otp string) (string, error)
	// VerifyOTP(ctx context.Context, phoneNumber string, otp string) (string, string, error)
	// RefreshToken(ctx context.Context, refreshToken string) (string, error)
	// DeleteRefreshToken(ctx context.Context, refreshToken string) error
}

type smsNotificationServiceGateway interface {
	Send(ctx context.Context, message string, from string, to string) (string, error)
}

type Controller struct {
	authenticationServiceGateway  authenticationServiceGateway
	smsNotificationServiceGateway smsNotificationServiceGateway
}

func New(
	authenticationServiceGateway authenticationServiceGateway,
	smsNotificationServiceGateway smsNotificationServiceGateway,
) *Controller {
	return &Controller{
		authenticationServiceGateway,
		smsNotificationServiceGateway,
	}
}

func (c *Controller) RequestOTP(ctx context.Context, phoneNumber string) (string, error) {
	otp, err := c.authenticationServiceGateway.GenerateOTP(ctx, phoneNumber)
	if err != nil {
		return "", err
	}
	message := fmt.Sprintf("Here is your 6-digit OTP: %s. Expires in 5mins", otp)
	sender := os.Getenv("SMSSENDER")
	sentResponse, err := c.smsNotificationServiceGateway.Send(ctx, message, sender, phoneNumber)
	if err != nil {
		return "", err
	}
	return sentResponse, nil
}
