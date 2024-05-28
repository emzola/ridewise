package grpc

import (
	"context"

	"github.com/emzola/ridewise/internal/grpcutil"
	"github.com/emzola/ridewise/pkg/discovery"
	pb "github.com/emzola/ridewise/smsnotificationservice/genproto"
)

type Gateway struct {
	registry discovery.Registry
}

func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry}
}

func (g *Gateway) Send(ctx context.Context, message, from, to string) (string, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "smsnotificationservice", g.registry)
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
