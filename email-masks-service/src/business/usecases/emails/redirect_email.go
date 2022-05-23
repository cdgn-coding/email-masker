package emails

import (
	"email-masks-service/src/business/gateways/emails"
	"email-masks-service/src/business/gateways/mappings"
)

type RedirectEmailUseCase interface {
	execute(email emails.Email) error
}

type redirectEmailUseCase struct {
	outboundEmailService emails.OutboundEmailService
	maskMappingService   mappings.MaskMappingService
}

func (r redirectEmailUseCase) execute(email emails.Email) error {
	ownerEmail, err := r.maskMappingService.GetOwnerEmail(email.To)
	if err != nil {
		return NewMaskAddressNotFoundError(email.To, err)
	}

	err = r.outboundEmailService.Send(emails.Email{
		To:          ownerEmail,
		Subject:     email.Subject,
		Content:     email.Content,
		ContentType: email.ContentType,
	})

	if err != nil {
		return NewOutboundEmailError(err)
	}

	return nil
}

func NewRedirectInboundEmail(outboundEmailService emails.OutboundEmailService, maskMappingService mappings.MaskMappingService) *redirectEmailUseCase {
	return &redirectEmailUseCase{
		outboundEmailService: outboundEmailService,
		maskMappingService:   maskMappingService,
	}
}
