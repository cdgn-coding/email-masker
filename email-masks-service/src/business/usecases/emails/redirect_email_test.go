package emails

import (
	"email-masks-service/src/business/gateways/emails"
	"errors"
	"fmt"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockedOutboundEmailService struct {
	mock.Mock
}

func (m *mockedOutboundEmailService) Send(email emails.Email) error {
	args := m.Called(email)
	return args.Error(0)
}

type mockedMaskMappingService struct {
	mock.Mock
}

func (m *mockedMaskMappingService) GetOwnerEmail(maskAddress string) (string, error) {
	args := m.Called(maskAddress)
	return args.String(0), args.Error(1)
}

func TestRedirectEmailUseCase(t *testing.T) {
	email := emails.Email{
		From:    "john@doe.com",
		To:      "mask@emailmasker.com",
		Subject: "",
		Content: "",
		HTML:    "",
	}

	t.Run("When mask is not found", func(t *testing.T) {
		emailService := new(mockedOutboundEmailService)
		maskMappingService := new(mockedMaskMappingService)
		redirectEmailUseCase := NewRedirectEmailUseCase(emailService, maskMappingService)

		maskMappingService.
			On("GetOwnerEmail", mock.Anything).
			Return("", fmt.Errorf("internal mask service error"))

		emailService.
			On("Send", mock.Anything).
			Return(nil)

		err := redirectEmailUseCase.Execute(email)

		if err == nil {
			t.Errorf("Expected to return an error")
			return
		}

		if !errors.Is(err, MaskAddressNotFoundError) {
			t.Errorf("Expected to return a MaskAddressNotFoundError. It was %v", err)
			return
		}
	})

	t.Run("When there is an error sending the email", func(t *testing.T) {
		emailService := new(mockedOutboundEmailService)
		maskMappingService := new(mockedMaskMappingService)
		redirectEmailUseCase := NewRedirectEmailUseCase(emailService, maskMappingService)

		maskMappingService.
			On("GetOwnerEmail", mock.Anything).
			Return("", nil)

		emailService.
			On("Send", mock.Anything).
			Return(fmt.Errorf("internal email service error"))

		err := redirectEmailUseCase.Execute(email)

		if err == nil {
			t.Errorf("Expected to return an error")
			return
		}

		if !errors.Is(err, OutboundEmailError) {
			t.Errorf("Expected to return a OutboundEmailError")
			return
		}
	})

	t.Run("When everything goes alright", func(t *testing.T) {
		emailService := new(mockedOutboundEmailService)
		maskMappingService := new(mockedMaskMappingService)
		redirectEmailUseCase := NewRedirectEmailUseCase(emailService, maskMappingService)

		maskMappingService.
			On("GetOwnerEmail", mock.Anything).
			Return("owner@email.com", nil)

		emailService.
			On("Send", mock.Anything).
			Return(nil)

		err := redirectEmailUseCase.Execute(email)

		if err != nil {
			t.Errorf("Expected not to return an error. %v was given", err)
			return
		}
	})
}
