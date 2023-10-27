package service

import (
	"errors"

	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/domain/user"
)

type UserService interface {
	CreateUser(username, password string) (*user.User, error)
}

type userService struct {
	repository user.Repository
}

func NewUserService(repository user.Repository) UserService {
	return &userService{
		repository: repository,
	}
}

func (s *userService) CreateUser(username, password string) (*user.User, error) {
	newUser := user.CreatBasicUser(username, password)

	exist, err := s.repository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if exist != nil {
		return nil, errors.New("user already exists")
	}

	err = s.repository.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return user.New(newUser.Username, "", "test-token", newUser.ID), nil
}
