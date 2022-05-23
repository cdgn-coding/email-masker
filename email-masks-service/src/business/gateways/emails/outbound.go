package emails

type Email struct {
	From        string
	To          string
	Subject     string
	Content     string
	ContentType string
}

type OutboundEmailService interface {
	Send(email Email) error
}
