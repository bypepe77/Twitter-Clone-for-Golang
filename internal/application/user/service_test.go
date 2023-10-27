package service

import (
	"errors"
	"testing"

	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/domain/user"
	jwtManager "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/jwt"
	repositories "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/repositories/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	defaultUser          = "user"
	defaultPassword      = "password"
	defaultUserID        = uint(0)
	errUserAlreadyExists = "user already exists"
)

func Test_Register_Happy_Path(t *testing.T) {
	repositoryMock := repositories.NewMockUserRepository(t)
	jwtManagerMock := jwtManager.NewMockManager(t)

	service := NewUserService(repositoryMock, jwtManagerMock)

	repositoryMock.EXPECT().GetUserByUsername(defaultUser).Return(nil, nil)
	jwtManagerMock.EXPECT().GenerateToken(defaultUserID, defaultUser).Return("token", nil)
	repositoryMock.EXPECT().CreateUser(mock.Anything).Return(nil)

	user, err := service.CreateUser(defaultUser, defaultPassword)
	assert.NoError(t, err)

	assert.Equal(t, defaultUser, user.Username)
	assert.Equal(t, "token", user.Token)
	assert.Equal(t, defaultUserID, user.ID)
}

func Test_Register_Should_Faild_Due_To_Existing_User(t *testing.T) {
	repositoryMock := repositories.NewMockUserRepository(t)
	jwtManagerMock := jwtManager.NewMockManager(t)

	service := NewUserService(repositoryMock, jwtManagerMock)

	repositoryMock.EXPECT().GetUserByUsername(defaultUser).Return(nil, errors.New(errUserAlreadyExists))

	_, err := service.CreateUser(defaultUser, defaultPassword)
	assert.Error(t, err)

	assert.Equal(t, errUserAlreadyExists, err.Error())
}

func Test_Login_Happy_Path(t *testing.T) {
	repositoryMock := repositories.NewMockUserRepository(t)
	jwtManagerMock := jwtManager.NewMockManager(t)

	service := NewUserService(repositoryMock, jwtManagerMock)
	hashedPassword, _ := hashPassword(defaultPassword)
	defaultUserForRepository := getDefaultUser(hashedPassword)

	repositoryMock.EXPECT().GetUserByUsername(defaultUser).Return(defaultUserForRepository, nil)
	jwtManagerMock.EXPECT().GenerateToken(defaultUserID, defaultUser).Return("token", nil)

	user, err := service.Login(defaultUser, defaultPassword)
	assert.NoError(t, err)

	assert.Equal(t, defaultUser, user.Username)
	assert.Equal(t, "token", user.Token)
	assert.Equal(t, defaultUserID, user.ID)
}

func Test_Login_Should_Fail_Due_To_Invalid_Credentials(t *testing.T) {
	repositoryMock := repositories.NewMockUserRepository(t)
	jwtManagerMock := jwtManager.NewMockManager(t)

	service := NewUserService(repositoryMock, jwtManagerMock)
	defaultUserForRepository := getDefaultUser(defaultPassword)

	repositoryMock.EXPECT().GetUserByUsername(defaultUser).Return(defaultUserForRepository, nil)

	_, err := service.Login(defaultUser, defaultPassword)
	assert.Error(t, err)
	assert.Equal(t, errInvalidCredentials, err.Error())
}

func Test_Login_Should_Fail_User_Not_Found(t *testing.T) {
	repositoryMock := repositories.NewMockUserRepository(t)
	jwtManagerMock := jwtManager.NewMockManager(t)

	service := NewUserService(repositoryMock, jwtManagerMock)

	repositoryMock.EXPECT().GetUserByUsername(defaultUser).Return(nil, nil)

	_, err := service.Login(defaultUser, defaultPassword)
	assert.Error(t, err)
	assert.Equal(t, errInvalidCredentials, err.Error())
}

func getDefaultUser(password string) *user.User {
	return &user.User{
		Username: defaultUser,
		Password: password,
	}
}
