package sendgrid

import (
	"email-masks-service/src/business/gateways"
	"errors"
	"github.com/sendgrid/sendgrid-go"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestOutboundEmailService_Send(t *testing.T) {
	t.Run("When SendGrid responds OK", func(t *testing.T) {
		defer gock.Off()

		gock.New("https://api.sendgrid.com/v3/mail/send").
			Post("/").
			Reply(200)

		client := sendgrid.NewSendClient("")
		emailService := NewOutboundEmailService(client)
		err := emailService.Send(gateways.Email{
			From:    "",
			To:      "",
			Subject: "",
			Content: "",
			HTML:    "",
		})

		if err != nil {
			t.Errorf("Expected not to return an error, given %v", err)
		}
	})

	t.Run("When SendGrid responds an error", func(t *testing.T) {
		defer gock.Off()

		gock.New("https://api.sendgrid.com/v3/mail/send").
			Post("/").
			Reply(500)

		client := sendgrid.NewSendClient("")
		emailService := NewOutboundEmailService(client)
		err := emailService.Send(gateways.Email{
			From:    "",
			To:      "",
			Subject: "",
			Content: "",
			HTML:    "",
		})

		if err == nil {
			t.Errorf("Expected to return an error")
			return
		}

		if !errors.Is(err, SendGridOutboundEmailError) {
			t.Errorf("Expected to return an SendGridOutboundEmailError, given %v", err)
		}
	})
}
