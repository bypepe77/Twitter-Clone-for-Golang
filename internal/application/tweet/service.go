package tweetservice

import (
	"context"

	tweetDomain "github.com/bypepe77/Twitter-Clone-for-Golang/internal/domain/tweet"
	jwtManager "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/jwt"
	"go.temporal.io/sdk/client"
)

type TweetService interface {
	CreateTweet(content string) error
}

type tweetService struct {
	jwtManager     jwtManager.Manager
	temporalClient client.Client
}

func New(jwtManager jwtManager.Manager, temporalClient client.Client) TweetService {
	return &tweetService{
		jwtManager:     jwtManager,
		temporalClient: temporalClient,
	}
}

func (s *tweetService) CreateTweet(content string) error {
	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: "TweetTaskQueue",
	}

	tweet := &tweetDomain.Tweet{
		Content:  content,
		UserID:   1,
		Username: "pepito",
	}

	_, err := s.temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, "ProcessTweetWorkflow", tweet)
	if err != nil {
		return err
	}

	return nil
}
