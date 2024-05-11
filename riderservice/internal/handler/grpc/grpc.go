package grpc

import (
	"context"
	"errors"

	pb "github.com/emzola/ridewise/genproto"
	"github.com/emzola/ridewise/riderservice/internal/controller"
	"github.com/emzola/ridewise/riderservice/internal/controller/rider"
	"github.com/emzola/ridewise/riderservice/pkg/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	pb.UnimplementedRiderServiceServer
	ctrl *rider.Controller
}

func New(ctrl *rider.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) CreateRider(ctx context.Context, req *pb.CreateRiderRequest) (*pb.CreateRiderResponse, error) {
	// TODO: Implement phone number validation
	if req == nil || req.Phone == "" {
		return nil, status.Errorf(codes.InvalidArgument, controller.ErrInvalidRequest.Error())
	}
	rider, err := h.ctrl.Create(ctx, req.Phone)
	if err != nil {
		code, errMsg := mapToGRPCErrorCode(err), err.Error()
		return &pb.CreateRiderResponse{
			Success:      false,
			Message:      errMsg,
			CreatedRider: nil,
		}, status.Errorf(code, errMsg)
	}
	return &pb.CreateRiderResponse{
		Success:      true,
		Message:      "Rider created successfully",
		CreatedRider: model.RiderToProto(rider),
	}, nil
}

func (h *Handler) GetRider(ctx context.Context, req *pb.GetRiderRequest) (*pb.GetRiderResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, controller.ErrInvalidRequest.Error())
	}
	rider, err := h.ctrl.Get(ctx, req.Id)
	if err != nil {
		code, errMsg := mapToGRPCErrorCode(err), err.Error()
		return nil, status.Errorf(code, errMsg)
	}
	return &pb.GetRiderResponse{Rider: model.RiderToProto(rider)}, nil
}

func (h *Handler) UpdateRider(ctx context.Context, req *pb.UpdateRiderRequest) (*pb.UpdateRiderResponse, error) {
	// TODO: Implement req data validation
	if req == nil || req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, controller.ErrInvalidRequest.Error())
	}
	// Convert the gRPC request message to the controller layer UpdateRiderRequest DTO struct.
	updateRequest := rider.UpdateRiderRequest{
		ID:             req.Id,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Phone:          req.Phone,
		Email:          req.Email,
		SavedLocations: convertSavedLocations(req.SavedLocations),
	}
	rider, err := h.ctrl.Update(ctx, updateRequest)
	if err != nil {
		code, errMsg := mapToGRPCErrorCode(err), err.Error()
		return &pb.UpdateRiderResponse{
			Success:      false,
			Message:      errMsg,
			UpdatedRider: nil,
		}, status.Errorf(code, errMsg)
	}
	return &pb.UpdateRiderResponse{
		Success:      true,
		Message:      "Rider updated successfully",
		UpdatedRider: model.RiderToProto(rider),
	}, nil
}

func (h *Handler) DeleteRider(ctx context.Context, req *pb.DeleteRiderRequest) (*pb.DeleteRiderResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, controller.ErrInvalidRequest.Error())
	}
	err := h.ctrl.Delete(ctx, req.Id)
	if err != nil {
		code, errMsg := mapToGRPCErrorCode(err), err.Error()
		return &pb.DeleteRiderResponse{
			Success: false,
			Message: errMsg,
		}, status.Errorf(code, errMsg)
	}
	return &pb.DeleteRiderResponse{
		Success: true,
		Message: "Rider deleted successfully",
	}, nil
}

// mapToGRPCErrorCode maps domain-specific errors to gRPC status codes.
func mapToGRPCErrorCode(err error) codes.Code {
	switch {
	case errors.Is(err, controller.ErrNotFound):
		return codes.NotFound
	case errors.Is(err, controller.ErrInvalidRequest):
		return codes.InvalidArgument
	case errors.Is(err, controller.ErrDuplicatePhone), errors.Is(err, controller.ErrDuplicateEmail):
		return codes.AlreadyExists
	default:
		return codes.Internal
	}
}

// convertSavedLocations converts saved locations from protobuf to the expected map type.
func convertSavedLocations(pbLocations map[string]*pb.Location) map[string]model.Location {
	savedLocations := map[string]model.Location{}
	for key, location := range pbLocations {
		savedLocations[key] = model.Location{
			Name:      location.Name,
			Latitude:  location.Latitude,
			Longitude: location.Longitude,
		}
	}
	return savedLocations
}