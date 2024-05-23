package httpsmsapi

import (
	"context"
	"os"

	"github.com/NdoleStudio/httpsms-go"
)

type Sender struct {
	client *httpsms.Client
}

func New() *Sender {
	return &Sender{client: httpsms.New(httpsms.WithAPIKey(os.Getenv("SMSAPIKEY")))}
}

func (s *Sender) Send(message string, from string, to string) error {
	_, _, err := s.client.Messages.Send(context.Background(), &httpsms.MessageSendParams{
		Content: message,
		From:    from,
		To:      to,
	})
	if err != nil {
		return err
	}
	return nil
}
