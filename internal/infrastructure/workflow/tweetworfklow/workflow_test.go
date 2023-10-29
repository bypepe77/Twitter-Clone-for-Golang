package tweetworfklow

import (
	"errors"
	"testing"

	tweetDomainModel "github.com/bypepe77/Twitter-Clone-for-Golang/internal/domain/tweet"
	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/testsuite"
)

func TestProcessTweetWorkflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	tweet := mockTweet()

	env.OnActivity(SaveTweetReference, tweet).Return(nil)

	env.ExecuteWorkflow(ProcessTweet, tweet)

	assert.True(t, env.IsWorkflowCompleted())
	assert.NoError(t, env.GetWorkflowError())

	env.AssertExpectations(t)
}

func TestProcessTweetWorkflow_Should_Fail(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	tweet := mockTweet()

	env.OnActivity(SaveTweetReference, tweet).Return(errors.New("error"))

	env.ExecuteWorkflow(ProcessTweet, tweet)

	assert.Error(t, env.GetWorkflowError())

	env.AssertExpectations(t)
}

func mockTweet() *tweetDomainModel.Tweet {
	return &tweetDomainModel.Tweet{
		Content:  "Hello World!",
		UserID:   1,
		Username: "pepito",
	}
}
