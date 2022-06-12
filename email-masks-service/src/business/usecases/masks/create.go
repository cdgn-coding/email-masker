package masks

import (
	"email-masks-service/src/business/entities"
	"email-masks-service/src/business/gateways"
	"email-masks-service/src/business/usecases"
	"fmt"
)

type ICreateMaskUseCase = usecases.CommandUseCase[*CreateMaskInput]

type CreateMaskUseCase struct {
	maskMappingService gateways.MaskMappingService
	usersService       gateways.UsersService
	maskDomain         string
}

func NewCreateMaskUseCase(
	maskMappingService gateways.MaskMappingService,
	usersService gateways.UsersService,
	maskDomain string,
) *CreateMaskUseCase {
	return &CreateMaskUseCase{maskMappingService: maskMappingService, usersService: usersService, maskDomain: maskDomain}
}

type CreateMaskInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
}

func (c *CreateMaskUseCase) createAddress(username, name string) string {
	return fmt.Sprintf("%s.%s@%s", username, name, c.maskDomain)
}

func (c *CreateMaskUseCase) Execute(input *CreateMaskInput) error {
	user, err := c.usersService.GetUserByID(input.UserID)
	if err != nil {
		return fmt.Errorf("%w. %v", FetchUserError, err)
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
		return fmt.Errorf("%w. %v", CreateMaskRecordError, err)
	}

	return nil
}
