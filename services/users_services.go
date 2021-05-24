package services

var (
	UsersService	userServiceInterface = &usersService{}
)

type usersService struct{}

type userServiceInterface interface {

}


