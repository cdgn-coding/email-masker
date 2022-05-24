package emails

import (
	"email-masks-service/src/business/gateways"
	"fmt"
)

type RedirectEmailUseCase interface {
	Execute(email gateways.Email) error
}

type redirectEmailUseCase struct {
	outboundEmailService gateways.OutboundEmailService
	maskMappingService   gateways.MaskMappingService
	usersService         gateways.UsersService
}

func (r redirectEmailUseCase) Execute(email gateways.Email) error {
	ownerUserID, err := r.maskMappingService.GetOwnerUserID(email.To)
	if err != nil {
		return fmt.Errorf("%v. %w", err, MaskAddressNotFoundError)
	}

	user, err := r.usersService.GetUserByID(ownerUserID)
	if err != nil {
		return fmt.Errorf("%v. %w", err, UserNotFoundError)
	}

	err = r.outboundEmailService.Send(gateways.Email{
		To:      user.Email,
		Subject: email.Subject,
		Content: email.Content,
		HTML:    email.HTML,
	})

	if err != nil {
		return fmt.Errorf("%v. %w", err, OutboundEmailError)
	}

	return nil
}

func NewRedirectEmailUseCase(
	outboundEmailService gateways.OutboundEmailService,
	maskMappingService gateways.MaskMappingService,
	usersService gateways.UsersService) *redirectEmailUseCase {
	return &redirectEmailUseCase{
		outboundEmailService: outboundEmailService,
		maskMappingService:   maskMappingService,
		usersService:         usersService,
	}
}
