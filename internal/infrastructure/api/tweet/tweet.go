package tweetapi

import (
	"net/http"

	tweetservice "github.com/bypepe77/Twitter-Clone-for-Golang/internal/application/tweet"
	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/api/responses"
	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/jwt"
	"github.com/gin-gonic/gin"
)

const (
	// ErrInvalidPayload is the error message when the payload is invalid
	ErrInvalidPayload = "invalid payload"
)

type TweetAPI interface {
	CreateTweet(c *gin.Context)
}

type tweetAPI struct {
	tweetService tweetservice.TweetService
	jwtManager   jwt.Manager
}

func New(tweetService tweetservice.TweetService, jwtManager jwt.Manager) TweetAPI {
	return &tweetAPI{
		tweetService: tweetService,
		jwtManager:   jwtManager,
	}
}

func (a *tweetAPI) CreateTweet(c *gin.Context) {
	var payload *tweetInput
	err := c.BindJSON(&payload)
	if err != nil {
		responses.Error(http.StatusBadRequest, c, ErrInvalidPayload)
		return
	}

	message, isValid := validateTweetInput(payload)
	if !isValid {
		responses.Error(http.StatusBadRequest, c, message)
		return
	}

	claims, err := a.jwtManager.GetClaims(c)
	if err != nil {
		responses.Error(http.StatusBadRequest, c, err.Error())
		return
	}

	err = a.tweetService.CreateTweet(payload.Content, claims)
	if err != nil {
		responses.Error(http.StatusBadRequest, c, err.Error())
		return
	}

	responses.Success(http.StatusCreated, c, "tweet created successfully")
}

func validateTweetInput(payload *tweetInput) (string, bool) {
	if payload.Content == "" {
		return "content is required", false
	}

	return "", true
}
