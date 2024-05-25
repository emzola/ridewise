package grpc

import (
	"context"

	pb "github.com/emzola/ridewise/smsnotificationservice/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Gateway struct{}

func New() *Gateway {
	return &Gateway{}
}

func (g *Gateway) Send(ctx context.Context, message, from, to string) (string, error) {
	conn, err := grpc.NewClient("localhost:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	defer conn.Close()
	client := pb.NewSMSNotificationServiceClient(conn)
	resp, err := client.SendSMS(ctx, &pb.SendSMSRequest{Content: message, From: from, To: to})
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}
