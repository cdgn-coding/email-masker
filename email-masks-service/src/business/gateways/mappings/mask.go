package mappings

type MaskMappingService interface {
	GetOwnerEmail(maskAddress string) (string, error)
}
