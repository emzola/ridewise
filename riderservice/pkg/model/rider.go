package model

type Rider struct {
	ID        string
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Places
}

type Places struct {
	Home       string
	Work       string
	Additional map[string]string
}
