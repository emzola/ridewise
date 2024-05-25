package grpc

import (
	"context"

	pb "github.com/emzola/ridewise/authenticationservice/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Gateway struct{}

func New() *Gateway {
	return &Gateway{}
}

func (g *Gateway) GenerateOTP(ctx context.Context, phoneNumber string) (string, error) {
	conn, err := grpc.NewClient("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
