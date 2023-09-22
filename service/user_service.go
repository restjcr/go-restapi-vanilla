package service

import (
	"go-restapi-v2/model"
	"go-restapi-v2/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) GetAllUsersService() (*[]model.User, error) {
	return s.userRepository.GetAllUsers()
}
