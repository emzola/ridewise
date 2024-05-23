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
	if req == nil || req.PhoneNumber == "" {
		return nil, status.Errorf(codes.InvalidArgument, controller.ErrInvalidRequest.Error())
	}
	otp, err := h.ctrl.GenerateOTP(ctx, req.PhoneNumber)
	if err != nil {
		code, errMsg := mapToGRPCErrorCode(err), err.Error()
		return nil, status.Errorf(code, errMsg)
	}
	return &pb.GenerateOTPResponse{Otp: otp}, nil
}

func (h *Handler) VerifyOTP(ctx context.Context, req *pb.VerifyOTPRequest) (*pb.VerifyOTPResponse, error) {
	if req == nil || req.Otp == "" {
		return nil, status.Errorf(codes.InvalidArgument, controller.ErrInvalidRequest.Error())
	}
	phoneNumber, err := h.ctrl.GetPhoneNumberByOTP(ctx, req.Otp)
	if err != nil {
		code, errMsg := mapToGRPCErrorCode(err), err.Error()
		return nil, status.Errorf(code, errMsg)
	}
	accessToken, refreshToken, err := h.ctrl.VerifyOTP(ctx, phoneNumber, req.Otp)
	if err != nil {
		code, errMsg := mapToGRPCErrorCode(err), err.Error()
		return nil, status.Errorf(code, errMsg)
	}
	return &pb.VerifyOTPResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (h *Handler) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	if req == nil || req.RefreshToken == "" {
		return nil, status.Errorf(codes.InvalidArgument, controller.ErrInvalidRequest.Error())
	}
	accessToken, err := h.ctrl.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		code, errMsg := mapToGRPCErrorCode(err), err.Error()
		return nil, status.Errorf(code, errMsg)
	}
	return &pb.RefreshTokenResponse{AccessToken: accessToken}, nil
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
