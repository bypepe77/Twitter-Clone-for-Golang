package tweetservice

import (
	"context"

	tweetDomain "github.com/bypepe77/Twitter-Clone-for-Golang/internal/domain/tweet"
	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/jwt"
	"go.temporal.io/sdk/client"
)

type TweetService interface {
	CreateTweet(content string, claims *jwt.Claims) error
}

type tweetService struct {
	temporalClient client.Client
}

func New(temporalClient client.Client) TweetService {
	return &tweetService{
		temporalClient: temporalClient,
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
