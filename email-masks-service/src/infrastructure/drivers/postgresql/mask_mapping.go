package postgresql

import (
	"email-masks-service/src/business/entities"
	"gorm.io/gorm"
)

type PostgresMaskMappingService struct {
	db *gorm.DB
}

func (p PostgresMaskMappingService) CreateMask(mask *entities.EmailMask) (*entities.EmailMask, error) {
	result := p.db.Create(mask)
	return mask, result.Error
}

func NewPostgresMaskMappingService(db *gorm.DB) *PostgresMaskMappingService {
	return &PostgresMaskMappingService{db: db}
}

func (p PostgresMaskMappingService) GetOwnerUserID(maskAddress string) (string, error) {
	emailMask := entities.EmailMask{}
	err := p.db.First(&emailMask, "address = ?", maskAddress).Error

	if err != nil {
		return "", err
	}

	return emailMask.UserID, nil
}
