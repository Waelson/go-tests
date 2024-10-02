package service

import (
	"github.com/Waelson/go-tests/internal/model"
	"github.com/Waelson/go-tests/internal/repository"
)

type UserService interface {
	FindAll() ([]model.User, error)
	Save(user model.User) (model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func (s *userService) FindAll() ([]model.User, error) {
	return s.userRepository.FindAll()
}

func (s *userService) Save(user model.User) (model.User, error) {
	return s.userRepository.Save(user)
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}
