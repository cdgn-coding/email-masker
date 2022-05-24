package auth0

import (
	"email-masks-service/src/business/entities"
	"github.com/auth0/go-auth0/management"
)

type Auth0UsersService struct {
	managementClient *management.Management
}

func NewAuth0UsersService(managementClient *management.Management) *Auth0UsersService {
	return &Auth0UsersService{managementClient: managementClient}
}

func (a Auth0UsersService) GetUserByID(userID string) (*entities.User, error) {
	auth0User, err := a.managementClient.User.Read(userID)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		ID:    *auth0User.ID,
		Email: *auth0User.Email,
	}

	return user, nil
}
