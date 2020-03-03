package emailSender

import (
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGrid struct {
	Key string
}

func (s *SendGrid) Send(to, from, subject, body string) error {
	fromEmail := mail.NewEmail("GDG Cloud Singapore", from)
	toEmail := mail.NewEmail("GDG Cloud Singapore member", to)
	message := mail.NewSingleEmail(fromEmail, subject, toEmail, body, body)
	client := sendgrid.NewSendClient(s.Key)
	response, err := client.Send(message)
	if err != nil {
		return err
	}
	if response.StatusCode != 202 {
		return fmt.Errorf("Wrong status code. %v", response.Body)
	}
	return nil
}
