package gateways

import "email-masks-service/src/business/entities"

type OutboundEmailService interface {
	Send(email entities.Email) error
}
