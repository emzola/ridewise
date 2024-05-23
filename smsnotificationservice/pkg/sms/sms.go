package sms

type Sender interface {
	Send(message string, from string, to string) error
}
