package services

import (
	"database/sql"
	"fmt"

	"github.com/go-playground/validator/v10"
	dto_user "github.com/mgambo/go-api/api/dto/user"
	"github.com/mgambo/go-api/api/models"
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

func (s *UserServiceImpl) CreateUser(user dto_user.CreateUserRequest) (*dto_user.UserResponse, error) {
	err := s.validate.Struct(user)
	if err != nil {
		return nil, err
	}

	userModel := &models.User{
		Username:    user.Username,
		Password:    []byte(user.Password),
		FirstName:   user.FirstName,
		LastName:    sql.NullString{String: user.LastName, Valid: true},
		DateOfBirth: user.DateOfBirth,
	}
	result, error := s.UserRepository.Base().Create(userModel)

	var response *dto_user.UserResponse

	response = &dto_user.UserResponse{
		Username:    result.Username,
		FirstName:   result.FirstName,
		LastName:    result.LastName.String,
		DateOfBirth: result.DateOfBirth,
	}

	return response, error
}

func (s *UserServiceImpl) GetUserById(id string) (*dto_user.UserResponse, error) {
	result, error := s.UserRepository.Base().GetByID(id)

	fmt.Println("abc", error)
	if result == nil {
		return nil, error
	}

	var response *dto_user.UserResponse

	response = &dto_user.UserResponse{
		Username:    result.Username,
		FirstName:   result.FirstName,
		LastName:    result.LastName.String,
		DateOfBirth: result.DateOfBirth,
	}

	return response, error
}
