package services

import (
	"github.com/go-playground/validator/v10"
	dto_user "github.com/mgambo/go-api/api/dto/user"
	"github.com/mgambo/go-api/api/repositories"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repositories.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		validate:       validate,
	}
}

// Implement the missing GetUsers method to satisfy the UserService interface.
func (s *UserServiceImpl) GetUsers() ([]dto_user.UserResponse, error) {
	result, error := s.UserRepository.Base().FindAll(nil)

	var response []dto_user.UserResponse

	for _, user := range result {
		response = append(response, dto_user.UserResponse{
			Username:    user.Username,
			FirstName:   user.FirstName,
			LastName:    user.LastName.String,
			DateOfBirth: user.DateOfBirth,
		})
	}

	return response, error
}
