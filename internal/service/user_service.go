package service

import (
	"github.com/Waelson/go-tests/internal/model"
	"github.com/Waelson/go-tests/internal/repository"
	log "github.com/sirupsen/logrus"
	"time"
)

type UserService interface {
	FindAll() ([]model.User, error)
	Save(user model.User) (model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func (s *userService) FindAll() ([]model.User, error) {
	start := time.Now()
	defer func() {
		log.Infof("[user_service] FindAll took %v", time.Since(start))
	}()
	return s.userRepository.FindAll()
}

func (s *userService) Save(user model.User) (model.User, error) {
	start := time.Now()
	defer func() {
		log.Infof("[user_service] Save took %v", time.Since(start))
	}()
	return s.userRepository.Save(user)
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}
