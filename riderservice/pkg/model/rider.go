package model

type Rider struct {
	ID             string
	FirstName      string
	LastName       string
	Phone          string
	Email          string
	IsVerified     bool
	SavedLocations map[string]Location
}

type Location struct {
	Name      string
	Latitude  float64
	Longitude float64
}
