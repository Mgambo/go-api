package services

import dto_user "github.com/mgambo/go-api/api/dto/user"

type UserService interface {
	GetUsers() ([]dto_user.UserResponse, error)
}
