package emails

import (
	"email-masks-service/src/business/gateways/emails"
	"email-masks-service/src/business/gateways/mappings"
	"fmt"
)

type RedirectEmailUseCase interface {
	Execute(email emails.Email) error
}

type redirectEmailUseCase struct {
	outboundEmailService emails.OutboundEmailService
	maskMappingService   mappings.MaskMappingService
}

func (r redirectEmailUseCase) Execute(email emails.Email) error {
	ownerEmail, err := r.maskMappingService.GetOwnerEmail(email.To)
	if err != nil {
		return fmt.Errorf("%v. %w", err, MaskAddressNotFoundError)
	}

	err = r.outboundEmailService.Send(emails.Email{
		To:      ownerEmail,
		Subject: email.Subject,
		Content: email.Content,
		HTML:    email.HTML,
	})

	if err != nil {
		return fmt.Errorf("%v. %w", err, OutboundEmailError)
	}

	return nil
}

func NewRedirectEmailUseCase(outboundEmailService emails.OutboundEmailService, maskMappingService mappings.MaskMappingService) *redirectEmailUseCase {
	return &redirectEmailUseCase{
		outboundEmailService: outboundEmailService,
		maskMappingService:   maskMappingService,
	}
}
