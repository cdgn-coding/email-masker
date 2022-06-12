package auth0

import (
	"email-masks-service/src/business/entities"
	"fmt"
	"github.com/auth0/go-auth0/management"
)

type Auth0UsersService struct {
	managementClient *management.Management
}

func NewAuth0UsersService(managementClient *management.Management) *Auth0UsersService {
	return &Auth0UsersService{managementClient: managementClient}
}

func (a Auth0UsersService) GetUserByID(userID string) (*entities.User, error) {
	formattedUserID := fmt.Sprintf("auth0|%s", userID)
	auth0User, err := a.managementClient.User.Read(formattedUserID)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		ID:       *auth0User.ID,
		Email:    *auth0User.Email,
		Username: *auth0User.Username,
	}

	return user, nil
}
