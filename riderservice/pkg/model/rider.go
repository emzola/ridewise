package model

type Rider struct {
	ID            string
	FirstName     string
	LastName      string
	Phone         string
	Activated     bool
	Email         string
	EmailVerified bool
	Places        Places
}

type Places struct {
	Home       string
	Work       string
	Additional map[string]string
}
