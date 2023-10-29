package tweetWorfklow

import (
	"time"

	tweetDomainModel "github.com/bypepe77/Twitter-Clone-for-Golang/internal/domain/tweet"
	"go.temporal.io/sdk/workflow"
)

type TweetWorkflow interface {
	ProcessTweet(ctx workflow.Context, tweet *tweetDomainModel.Tweet) error
}

type tweetWorkflow struct {
	activities TweetActivities
}

func NewTweetWorkflow(activities TweetActivities) TweetWorkflow {
	return &tweetWorkflow{activities: activities}
}

func (w *tweetWorkflow) ProcessTweet(ctx workflow.Context, tweet *tweetDomainModel.Tweet) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, w.activities.SaveTweet, tweet).Get(ctx, nil)
	return err
}
