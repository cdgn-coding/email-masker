package gateways

import "email-masks-service/src/business/entities"

type MaskMappingService interface {
	GetOwnerUserID(maskAddress string) (string, error)
	CreateMask(mask *entities.EmailMask) (entities.EmailMask, error)
}
