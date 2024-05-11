package grpc

import (
	"context"
	"errors"

	"github.com/emzola/ridewise/gen"
	"github.com/emzola/ridewise/riderservice/internal/controller"
	"github.com/emzola/ridewise/riderservice/internal/controller/rider"
	"github.com/emzola/ridewise/riderservice/pkg/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	gen.UnimplementedRiderServiceServer
	ctrl *rider.Controller
}

func New(ctrl *rider.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) CreateRider(ctx context.Context, req *gen.CreateRiderRequest) (*gen.CreateRiderResponse, error) {
	if req == nil || req.Phone == "" {
		return nil, status.Errorf(codes.InvalidArgument, controller.ErrInvalidRequest.Error())
	}
	rider, err := h.ctrl.Create(ctx, req.Phone)
	if err != nil {
		code, errMsg := mapToGRPCErrorCode(err), err.Error()
		return &gen.CreateRiderResponse{
			Success:      false,
			Message:      errMsg,
			CreatedRider: nil,
		}, status.Errorf(code, errMsg)
	}
	return &gen.CreateRiderResponse{
		Success:      true,
		Message:      "Rider created successfully",
		CreatedRider: model.RiderToProto(rider),
	}, nil
}

func (h *Handler) GetRider(ctx context.Context, req *gen.GetRiderRequest) (*gen.GetRiderResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, controller.ErrInvalidRequest.Error())
	}
	rider, err := h.ctrl.Get(ctx, req.Id)
	if err != nil {
		code, errMsg := mapToGRPCErrorCode(err), err.Error()
		return nil, status.Errorf(code, errMsg)
	}
	return &gen.GetRiderResponse{Rider: model.RiderToProto(rider)}, nil
}

func (h *Handler) UpdateRider(ctx context.Context, req *gen.UpdateRiderRequest) (*gen.UpdateRiderResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, controller.ErrInvalidRequest.Error())
	}
	rider, err := h.ctrl.Update(ctx, req)
	if err != nil {
		code, errMsg := mapToGRPCErrorCode(err), err.Error()
		return &gen.UpdateRiderResponse{
			Success:      false,
			Message:      errMsg,
			UpdatedRider: nil,
		}, status.Errorf(code, errMsg)
	}
	return &gen.UpdateRiderResponse{
		Success:      true,
		Message:      "Rider updated successfully",
		UpdatedRider: model.RiderToProto(rider),
	}, nil
}

func (h *Handler) DeleteRider(ctx context.Context, req *gen.DeleteRiderRequest) (*gen.DeleteRiderResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, controller.ErrInvalidRequest.Error())
	}
	err := h.ctrl.Delete(ctx, req.Id)
	if err != nil {
		code, errMsg := mapToGRPCErrorCode(err), err.Error()
		return &gen.DeleteRiderResponse{
			Success: false,
			Message: errMsg,
		}, status.Errorf(code, errMsg)
	}
	return &gen.DeleteRiderResponse{
		Success: true,
		Message: "Rider deleted successfully",
	}, nil
}

// MapToGRPCErrorCode maps domain-specific errors to gRPC status codes.
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
