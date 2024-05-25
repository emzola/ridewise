package grpc

import (
	"context"
	"errors"

	pb "github.com/emzola/ridewise/gatewayservice/genproto"
	"github.com/emzola/ridewise/gatewayservice/internal/controller"
	"github.com/emzola/ridewise/gatewayservice/internal/controller/gateway"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	pb.UnimplementedGatewayServiceServer
	ctrl *gateway.Controller
}

func New(ctrl *gateway.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) RequestOTP(ctx context.Context, req *pb.RequestOTPRequest) (*pb.RequestOTPResponse, error) {
	if req == nil || req.PhoneNumber == "" {
		return nil, status.Errorf(codes.InvalidArgument, controller.ErrInvalidRequest.Error())
	}
	sentResponse, err := h.ctrl.RequestOTP(ctx, req.PhoneNumber)
	if err != nil {
		code, errMsg := mapToGRPCErrorCode(err), err.Error()
		return nil, status.Errorf(code, errMsg)
	}
	return &pb.RequestOTPResponse{Message: sentResponse}, nil
}

// mapToGRPCErrorCode maps domain-specific errors to gRPC status codes.
func mapToGRPCErrorCode(err error) codes.Code {
	switch {
	case errors.Is(err, controller.ErrNotFound):
		return codes.NotFound
	case errors.Is(err, controller.ErrInvalidRequest), errors.Is(err, controller.ErrInvalidOTP), errors.Is(err, controller.ErrInvalidRefreshToken):
		return codes.InvalidArgument
	default:
		return codes.Internal
	}
}
