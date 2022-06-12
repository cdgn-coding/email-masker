package gateways

import "email-masks-service/src/business/entities"

type UsersService interface {
	GetUserByID(userID string) (*entities.User, error)
}
