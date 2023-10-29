package tweetWorfklow

import (
	tweetDomainModel "github.com/bypepe77/Twitter-Clone-for-Golang/internal/domain/tweet"
	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/repositories/tweet"
)

var SaveTweetReference = (&tweetActivities{}).SaveTweet

type TweetActivities interface {
	SaveTweet(tweet *tweetDomainModel.Tweet) error
}

type tweetActivities struct {
	tweetRepository tweet.Repository
}

func NewTweetActivities(tweetRepository tweet.Repository) TweetActivities {
	return &tweetActivities{
		tweetRepository: tweetRepository,
	}
}

// SaveTweet saves a tweet in the database
func (a *tweetActivities) SaveTweet(tweet *tweetDomainModel.Tweet) error {
	return a.tweetRepository.SaveTweet(tweet)
}
