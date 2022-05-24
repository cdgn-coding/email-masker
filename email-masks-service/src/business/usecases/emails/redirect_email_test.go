package emails

import (
	"email-masks-service/src/business/entities"
	"email-masks-service/src/business/gateways"
	"errors"
	"fmt"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockedOutboundEmailService struct {
	mock.Mock
}

func (m *mockedOutboundEmailService) Send(email gateways.Email) error {
	args := m.Called(email)
	return args.Error(0)
}

type mockedMaskMappingService struct {
	mock.Mock
}

func (m *mockedMaskMappingService) GetOwnerUserID(maskAddress string) (string, error) {
	args := m.Called(maskAddress)
	return args.String(0), args.Error(1)
}

type mockedUsersService struct {
	mock.Mock
}

func (m *mockedUsersService) GetUserByID(userID string) (*entities.User, error) {
	args := m.Called(userID)
	return args.Get(0).(*entities.User), args.Error(1)
}

func TestRedirectEmailUseCase_Execute(t *testing.T) {
	email := gateways.Email{
		From:    "john@doe.com",
		To:      "mask@emailmasker.com",
		Subject: "",
		Content: "",
		HTML:    "",
	}

	t.Run("When mask is not found", func(t *testing.T) {
		emailService := new(mockedOutboundEmailService)
		maskMappingService := new(mockedMaskMappingService)
		usersService := new(mockedUsersService)
		redirectEmailUseCase := NewRedirectEmailUseCase(emailService, maskMappingService, usersService)

		maskMappingService.
			On("GetOwnerUserID", mock.Anything).
			Return("", fmt.Errorf("internal mask service error"))

		usersService.
			On("GetUserByID", mock.Anything).
			Return(&entities.User{}, fmt.Errorf("internal users service error"))

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
		usersService := new(mockedUsersService)
		redirectEmailUseCase := NewRedirectEmailUseCase(emailService, maskMappingService, usersService)

		maskMappingService.
			On("GetOwnerUserID", mock.Anything).
			Return("longUserID", nil)

		usersService.
			On("GetUserByID", "longUserID").
			Return(&entities.User{Email: "anEmail"}, nil)

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

	t.Run("When there is an error obtaining the user", func(t *testing.T) {
		emailService := new(mockedOutboundEmailService)
		maskMappingService := new(mockedMaskMappingService)
		usersService := new(mockedUsersService)
		redirectEmailUseCase := NewRedirectEmailUseCase(emailService, maskMappingService, usersService)

		maskMappingService.
			On("GetOwnerUserID", mock.Anything).
			Return("longUserID", nil)

		usersService.
			On("GetUserByID", "longUserID").
			Return(&entities.User{}, fmt.Errorf("internal users service error"))

		emailService.
			On("Send", mock.Anything).
			Return(nil)

		err := redirectEmailUseCase.Execute(email)

		if err == nil {
			t.Errorf("Expected not to return an error.")
			return
		}

		if !errors.Is(err, UserNotFoundError) {
			t.Errorf("Expected not to return a UserNotFoundError, given %v", err)
			return
		}
	})

	t.Run("When everything goes alright", func(t *testing.T) {
		emailService := new(mockedOutboundEmailService)
		maskMappingService := new(mockedMaskMappingService)
		usersService := new(mockedUsersService)
		redirectEmailUseCase := NewRedirectEmailUseCase(emailService, maskMappingService, usersService)

		maskMappingService.
			On("GetOwnerUserID", mock.Anything).
			Return("longUserID", nil)

		usersService.
			On("GetUserByID", "longUserID").
			Return(&entities.User{Email: "anEmail"}, nil)

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
