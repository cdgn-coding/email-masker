package masks

import (
	"email-masks-service/src/business/entities"
	"email-masks-service/src/business/gateways"
	"fmt"
)

type CreateMaskUseCase struct {
	maskMappingService gateways.MaskMappingService
	usersService       gateways.UsersService
	maskDomain         string
}

type CreateMaskInput struct {
	Address     string `json:"address" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
}

func (c *CreateMaskUseCase) createAddress(username, name string) string {
	return fmt.Sprintf("%s.%s@%s", username, name, c.maskDomain)
}

func (c *CreateMaskUseCase) Execute(input CreateMaskInput) error {
	user, err := c.usersService.GetUserByID(input.UserID)
	if err != nil {
		return fmt.Errorf("%v. %w", err, ErrorCreatingMask)
	}

	address := c.createAddress(user.Username, input.Name)
	mask := &entities.EmailMask{
		Address: address,
		Name:    input.Name,
		UserID:  input.UserID,
		Enabled: true,
	}

	_, err = c.maskMappingService.CreateMask(mask)
	if err != nil {
		return fmt.Errorf("%v. %w", err, ErrorCreatingMask)
	}

	return nil
}
