package tweetworfklow

import (
	"time"

	tweetDomainModel "github.com/bypepe77/Twitter-Clone-for-Golang/internal/domain/tweet"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func ProcessTweet(ctx workflow.Context, tweet *tweetDomainModel.Tweet) error {
	ctx = withActivityOptions(ctx)

	err := workflow.ExecuteActivity(ctx, SaveTweetReference, tweet).Get(ctx, nil)
	return err
}

func withActivityOptions(ctx workflow.Context) workflow.Context {
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Minute,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Minute * 32,
		MaximumAttempts:    3,
	}
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retrypolicy,
	}

	return workflow.WithActivityOptions(ctx, options)
}
