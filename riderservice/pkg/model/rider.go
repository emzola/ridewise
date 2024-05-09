package model

type Rider struct {
	ID        string
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Addresses
}

type Addresses struct {
	Home       string
	Work       string
	Additional map[string]string
}
