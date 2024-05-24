package model

import (
	pb "github.com/emzola/ridewise/riderservice/genproto"
)

// RiderToProto converts a Rider struct into a generated proto counterpart.
func RiderToProto(rider *Rider) *pb.Rider {
	riderProto := &pb.Rider{
		Id:             rider.ID,
		FirstName:      rider.FirstName,
		LastName:       rider.LastName,
		Phone:          rider.Phone,
		Email:          rider.Email,
		IsVerified:     rider.IsVerified,
		SavedLocations: map[string]*pb.Location{},
	}
	// Convert saved locations
	for key, location := range rider.SavedLocations {
		riderProto.SavedLocations[key] = &pb.Location{
			Name:      location.Name,
			Latitude:  location.Latitude,
			Longitude: location.Longitude,
		}
	}
	return riderProto
}

// RiderFromProto converts a generated proto counterpart into a Rider struct.
func RiderFromProto(riderProto *pb.Rider) *Rider {
	rider := &Rider{
		ID:             riderProto.Id,
		FirstName:      riderProto.FirstName,
		LastName:       riderProto.LastName,
		Phone:          riderProto.Phone,
		Email:          riderProto.Email,
		IsVerified:     riderProto.IsVerified,
		SavedLocations: map[string]Location{},
	}
	// Convert saved locations
	for key, location := range riderProto.SavedLocations {
		rider.SavedLocations[key] = Location{
			Name:      location.Name,
			Latitude:  location.Latitude,
			Longitude: location.Longitude,
		}
	}
	return rider
}
