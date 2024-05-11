package model

import (
	"github.com/emzola/ridewise/gen"
)

// RiderToProto converts a Rider struct into a generated proto counterpart.
func RiderToProto(r *Rider) *gen.Rider {
	return &gen.Rider{
		Id:        r.ID,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Phone:     r.Phone,
		Email:     r.Email,
		Places: &gen.Places{
			Home:       r.Places.Home,
			Work:       r.Places.Work,
			Additional: r.Places.Additional,
		},
	}
}

// RiderFromProto converts a generated proto counterpart into a Rider struct.
func RiderFromProto(r *gen.Rider) *Rider {
	return &Rider{
		ID:        r.Id,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Phone:     r.Phone,
		Email:     r.Email,
		Places: Places{
			Home:       r.Places.Home,
			Work:       r.Places.Work,
			Additional: r.Places.Additional,
		},
	}
}
