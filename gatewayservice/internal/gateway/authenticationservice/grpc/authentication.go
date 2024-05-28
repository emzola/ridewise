package grpc

import (
	"context"

	pb "github.com/emzola/ridewise/authenticationservice/genproto"
	"github.com/emzola/ridewise/internal/grpcutil"
	"github.com/emzola/ridewise/pkg/discovery"
)

type Gateway struct {
	registry discovery.Registry
}

func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry}
}

func (g *Gateway) GenerateOTP(ctx context.Context, phoneNumber string) (string, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "authenticationservice", g.registry)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	client := pb.NewAuthenticationServiceClient(conn)
	resp, err := client.GenerateOTP(ctx, &pb.GenerateOTPRequest{PhoneNumber: phoneNumber})
	if err != nil {
		return "", err
	}
	return resp.Otp, nil
}
