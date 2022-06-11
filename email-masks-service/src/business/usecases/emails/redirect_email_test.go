package emails

import (
	"email-masks-service/src/business/entities"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockedOutboundEmailService struct {
	mock.Mock
}

func (m *mockedOutboundEmailService) Send(email entities.Email) error {
	args := m.Called(email)
	return args.Error(0)
}

type mockedMaskMappingService struct {
	mock.Mock
}

func (m *mockedMaskMappingService) CreateMask(mask *entities.EmailMask) (*entities.EmailMask, error) {
	panic("not implemented")
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
	email := &entities.Email{
		From:    "john@doe.com",
		To:      "mask@emailmasker.com",
		Subject: "",
		Content: "",
		HTML:    "",
	}

	t.Run("mask not found error", func(t *testing.T) {
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

		assert.Error(t, err)
		assert.ErrorIs(t, err, MaskAddressNotFoundError)
	})

	t.Run("outbound email error", func(t *testing.T) {
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

		assert.Error(t, err)
		assert.ErrorIs(t, err, OutboundEmailError)
	})

	t.Run("user not found error", func(t *testing.T) {
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

		assert.Error(t, err)
		assert.ErrorIs(t, err, UserNotFoundError)
	})

	t.Run("success", func(t *testing.T) {
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

		assert.NoError(t, err)
	})
}
