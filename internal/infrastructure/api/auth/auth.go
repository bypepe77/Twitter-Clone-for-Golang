package auth

import (
	"net/http"

	service "github.com/bypepe77/Twitter-Clone-for-Golang/internal/application/user"
	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/api/responses"
	"github.com/gin-gonic/gin"
)

type Auth interface {
	CreateUser(c *gin.Context)
}

type auth struct {
	service service.UserService
}

func New(service service.UserService) Auth {
	return &auth{
		service: service,
	}
}

func (a *auth) CreateUser(c *gin.Context) {
	var payload *input
	err := c.BindJSON(&payload)
	if err != nil {
		responses.Error(http.StatusBadRequest, c, ErrInvalidPayload)
		return
	}

	message, isValid := validateInput(payload)
	if !isValid {
		responses.Error(http.StatusBadRequest, c, message)
		return
	}

	user, err := a.service.CreateUser(payload.Username, payload.Password)
	if err != nil {
		responses.Error(http.StatusBadRequest, c, err.Error())
		return
	}

	userCreated := NewUserResponse(user.Token, user.Username, user.ID)
	responses.Success(http.StatusCreated, c, userCreated)
}

func validateInput(input *input) (string, bool) {
	if input.Username == "" {
		return ErrEmptyUsername, false
	}

	if input.Password == "" {
		return ErrEmptyPassword, false
	}

	return "", true
}
