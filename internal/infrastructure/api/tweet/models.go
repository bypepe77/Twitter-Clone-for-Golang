package tweetapi

import "github.com/bypepe77/Twitter-Clone-for-Golang/internal/domain/tweet"

type tweetInput struct {
	Content string `json:"content"`
}

type TweetResponse struct {
	ID            uint   `json:"id"`
	Content       string `json:"content"`
	UserID        uint   `json:"user_id"`
	Username      string `json:"username"`
	LikesCount    int    `json:"likes_count"`
	RetweetsCount int    `json:"retweets_count"`
	RepliesCount  int    `json:"replies_count"`
}

func toTweetResponse(tweet *tweet.Tweet) *TweetResponse {
	return &TweetResponse{
		ID:            tweet.ID,
		Content:       tweet.Content,
		UserID:        tweet.UserID,
		Username:      tweet.Username,
		LikesCount:    tweet.LikesCount,
		RetweetsCount: tweet.RetweetsCount,
		RepliesCount:  tweet.RepliesCount,
	}
}
