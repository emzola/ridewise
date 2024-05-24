package grpc

import (
	"context"
	"errors"

	pb "github.com/emzola/ridewise/smsnotificationservice/genproto"
	"github.com/emzola/ridewise/smsnotificationservice/internal/controller"
	"github.com/emzola/ridewise/smsnotificationservice/internal/controller/sms"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	ctrl *sms.Controller
	pb.UnimplementedSMSNotificationServiceServer
}

func New(ctrl *sms.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) SendSMS(ctx context.Context, req *pb.SendSMSRequest) (*pb.SendSMSResponse, error) {
	if req == nil || req.Content == "" || req.From == "" || req.To == "" {
		return nil, status.Errorf(codes.InvalidArgument, controller.ErrInvalidRequest.Error())
	}
	err := h.ctrl.Send(ctx, req.Content, req.From, req.To)
	if err != nil {
		code, errMsg := mapToGRPCErrorCode(err), err.Error()
		return nil, status.Errorf(code, errMsg)
	}
	return &pb.SendSMSResponse{Message: "sms sent"}, nil
}

// mapToGRPCErrorCode maps domain-specific errors to gRPC status codes.
func mapToGRPCErrorCode(err error) codes.Code {
	switch {
	case errors.Is(err, controller.ErrInvalidRequest):
		return codes.InvalidArgument
	default:
		return codes.Internal
	}
}
