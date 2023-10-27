package service

import (
	"errors"

	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/domain/user"
	jwtManager "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/jwt"
	"golang.org/x/crypto/bcrypt"
)

const (
	// ErrInvalidCredentials is returned when the credentials are invalid.
	errInvalidCredentials = "invalid credentials"
)

type UserService interface {
	CreateUser(username, password string) (*user.User, error)
	Login(username, password string) (*user.User, error)
}

type userService struct {
	repository user.Repository
	jwtManager jwtManager.Manager
}

func NewUserService(repository user.Repository, jwtManager jwtManager.Manager) UserService {
	return &userService{
		repository: repository,
		jwtManager: jwtManager,
	}
}

func (s *userService) CreateUser(username, password string) (*user.User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	newUser := user.CreatBasicUser(username, hashedPassword)

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

	token, err := s.jwtManager.GenerateToken(newUser.ID, newUser.Username)
	if err != nil {
		return nil, err
	}

	return user.New(newUser.Username, "", token, newUser.ID), nil
}

func (s *userService) Login(username, password string) (*user.User, error) {
	userRetrieved, err := s.repository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if userRetrieved == nil {
		return nil, errors.New(errInvalidCredentials)
	}

	if !checkPasswordHash(password, userRetrieved.Password) {
		return nil, errors.New(errInvalidCredentials)
	}

	token, err := s.jwtManager.GenerateToken(userRetrieved.ID, userRetrieved.Username)
	if err != nil {
		return nil, err
	}

	return user.New(userRetrieved.Username, "", token, userRetrieved.ID), nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

//nolint:deadcode
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
