package grpc

import (
	"context"
	"errors"

	"github.com/emzola/ridewise/authenticationservice/internal/controller"
	"github.com/emzola/ridewise/authenticationservice/internal/controller/auth"
	pb "github.com/emzola/ridewise/genproto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	ctrl *auth.Controller
	pb.UnimplementedAuthenticationServiceServer
}

func New(ctrl *auth.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) GenerateOTP(ctx context.Context, req *pb.GenerateOTPRequest) (*pb.GenerateOTPResponse, error) {
	if req == nil || req.Phone == "" {
		return nil, status.Errorf(codes.InvalidArgument, controller.ErrInvalidRequest.Error())
	}
	otp, err := h.ctrl.GenerateOTP(req.Phone)
	if err != nil {
		code, errMsg := mapToGRPCErrorCode(err), err.Error()
		return nil, status.Errorf(code, errMsg)
	}
	return &pb.GenerateOTPResponse{Otp: otp}, nil
}

// mapToGRPCErrorCode maps domain-specific errors to gRPC status codes.
func mapToGRPCErrorCode(err error) codes.Code {
	switch {
	case errors.Is(err, controller.ErrNotFound):
		return codes.NotFound
	case errors.Is(err, controller.ErrInvalidRequest):
		return codes.InvalidArgument
	default:
		return codes.Internal
	}
}
