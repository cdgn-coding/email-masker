package sendgrid

import (
	"email-masks-service/src/business/entities"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"net/http"
)

type OutboundEmailService struct {
	client *sendgrid.Client
}

func NewOutboundEmailService(client *sendgrid.Client) *OutboundEmailService {
	return &OutboundEmailService{
		client: client,
	}
}

func (s *OutboundEmailService) Send(email entities.Email) error {
	from := mail.NewEmail("", email.From)
	to := mail.NewEmail("", email.To)
	message := mail.NewSingleEmail(from, email.Subject, to, email.Content, email.HTML)
	resp, err := s.client.Send(message)

	if err != nil {
		return fmt.Errorf("%w. %v", SendGridOutboundEmailError, err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w. Status Code %d", SendGridOutboundEmailError, resp.StatusCode)
	}

	return nil
}
