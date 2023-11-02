package tweetservice

import (
	"context"
	"errors"

	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/domain/tweet"
	tweetDomain "github.com/bypepe77/Twitter-Clone-for-Golang/internal/domain/tweet"
	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/jwt"
	tweetRepository "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/repositories/tweet"
	"go.temporal.io/sdk/client"
)

type TweetService interface {
	CreateTweet(content string, claims *jwt.Claims) error
	GetTweet(tweetID uint) (*tweet.Tweet, error)
}

type tweetService struct {
	temporalClient  client.Client
	tweetRepository tweetRepository.Repository
}

func New(temporalClient client.Client, tweetRepository tweetRepository.Repository) TweetService {
	return &tweetService{
		temporalClient:  temporalClient,
		tweetRepository: tweetRepository,
	}
}

func (s *tweetService) CreateTweet(content string, claims *jwt.Claims) error {
	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: "TweetTaskQueue",
	}

	tweet := &tweetDomain.Tweet{
		Content:  content,
		UserID:   claims.UserID,
		Username: claims.Username,
	}

	_, err := s.temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, "ProcessTweetWorkflow", tweet)
	if err != nil {
		return err
	}

	return nil
}

func (s *tweetService) GetTweet(tweetID uint) (*tweet.Tweet, error) {
	if tweetID == 0 {
		return nil, errors.New("invalid tweet id")
	}

	tweet, err := s.tweetRepository.GetTweetByID(tweetID)
	if err != nil {
		return nil, err
	}

	return tweet, nil
}
