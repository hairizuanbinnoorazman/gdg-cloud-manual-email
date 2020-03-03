package emailSender

type EmailSender interface {
	Send(to, from, subject, body string) error
}
