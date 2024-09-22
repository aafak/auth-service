package service

// Service contains actual business logic

import (
	"github.com/aafak/auth-service/internal/model"
	"github.com/aafak/auth-service/internal/repository"
)

type UserService interface {
	Create(user *model.User) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepository: userRepo}
}

func (s *userService) Create(user *model.User) error {
	return s.userRepository.Create(user)
}
