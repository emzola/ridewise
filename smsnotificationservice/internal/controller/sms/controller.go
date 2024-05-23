package sms

import (
	"context"

	"github.com/emzola/ridewise/smsnotificationservice/pkg/sms"
)

type Controller struct {
	smsAPI sms.Sender
}

func New(smsAPI sms.Sender) *Controller {
	return &Controller{smsAPI}
}

func (c *Controller) Send(ctx context.Context, message string, from string, to string) error {
	return c.smsAPI.Send(message, from, to)
}
