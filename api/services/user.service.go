package services

import dto_user "github.com/mgambo/go-api/api/dto/user"

type UserService interface {
	GetUsers() ([]dto_user.UserResponse, error)
	GetUserById(id string) (*dto_user.UserResponse, error)
	CreateUser(user dto_user.CreateUserRequest) (*dto_user.UserResponse, error)
}
