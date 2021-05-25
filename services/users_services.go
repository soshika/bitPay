package services

import (
	"bitPay/domains/users"
	"bitPay/logger"
	"bitPay/utills/responses"
)

var (
	UsersService	userServiceInterface = &usersService{}
)

type usersService struct{}

type userServiceInterface interface {
	Register(users.User) *responses.Response
}

func(s *usersService) Register(user users.User) *responses.Response {
	logger.Info("Enter to Register service successfully")

	if saveErr := user.Save(); saveErr != nil {
		return saveErr
	}

	logger.Info("Close from Register service successfully")
	return nil
}


