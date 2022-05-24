package gateways

type Email struct {
	From    string
	To      string
	Subject string
	Content string
	HTML    string
}

type OutboundEmailService interface {
	Send(email Email) error
}
