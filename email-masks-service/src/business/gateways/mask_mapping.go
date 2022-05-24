package gateways

type MaskMappingService interface {
	GetOwnerUserID(maskAddress string) (string, error)
}
