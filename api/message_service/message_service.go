package userserviceimpl

import (
	"github.com/AliMumtazDev/Go_Chat_App/database/mongodb"
)

type UserServiceImpl struct {
	messageAuth mongodb.Storage
}

func NewUserService(input mongodb.Storage) UserService {
	return &UserServiceImpl{
		messageAuth: input,
	}
}

type NewUserServiceImpl struct {
	messageAuth mongodb.Storage
}

var _ UserService = &UserServiceImpl{}
