package sendgrid

import (
	"email-masks-service/src/business/entities"
	"github.com/sendgrid/sendgrid-go"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestOutboundEmailService_Send(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		defer gock.Off()

		gock.New("https://api.sendgrid.com/v3/mail/send").
			Post("/").
			Reply(200)

		client := sendgrid.NewSendClient("")
		emailService := NewOutboundEmailService(client)
		err := emailService.Send(entities.Email{
			From:    "",
			To:      "",
			Subject: "",
			Content: "",
			HTML:    "",
		})

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		defer gock.Off()

		gock.New("https://api.sendgrid.com/v3/mail/send").
			Post("/").
			Reply(500)

		client := sendgrid.NewSendClient("")
		emailService := NewOutboundEmailService(client)
		err := emailService.Send(entities.Email{
			From:    "",
			To:      "",
			Subject: "",
			Content: "",
			HTML:    "",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, SendGridOutboundEmailError)
	})
}
