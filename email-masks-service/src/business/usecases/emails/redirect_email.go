package emails

import (
	"email-masks-service/src/business/entities"
	"email-masks-service/src/business/gateways"
	"email-masks-service/src/business/usecases"
	"fmt"
)

type IRedirectEmailUseCase = usecases.CommandUseCase[*entities.Email]

type RedirectEmailUseCase struct {
	outboundEmailService gateways.OutboundEmailService
	maskMappingService   gateways.MaskMappingService
	usersService         gateways.UsersService
}

func (r RedirectEmailUseCase) Execute(email *entities.Email) error {
	ownerUserID, err := r.maskMappingService.GetOwnerUserID(email.To)
	if err != nil {
		return fmt.Errorf("%w. %v", MaskAddressNotFoundError, err)
	}

	user, err := r.usersService.GetUserByID(ownerUserID)
	if err != nil {
		return fmt.Errorf("%w. %v", UserNotFoundError, err)
	}

	err = r.outboundEmailService.Send(entities.Email{
		To:      user.Email,
		Subject: email.Subject,
		Content: email.Content,
		HTML:    email.HTML,
	})

	if err != nil {
		return fmt.Errorf("%w. %v", OutboundEmailError, err)
	}

	return nil
}

func NewRedirectEmailUseCase(
	outboundEmailService gateways.OutboundEmailService,
	maskMappingService gateways.MaskMappingService,
	usersService gateways.UsersService) *RedirectEmailUseCase {
	return &RedirectEmailUseCase{
		outboundEmailService: outboundEmailService,
		maskMappingService:   maskMappingService,
		usersService:         usersService,
	}
}
