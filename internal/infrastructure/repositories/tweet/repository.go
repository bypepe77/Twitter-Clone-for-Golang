package tweet

import (
	tweetDomainModel "github.com/bypepe77/Twitter-Clone-for-Golang/internal/domain/tweet"
	"gorm.io/gorm"
)

type Repository interface {
	SaveTweet(tweet *tweetDomainModel.Tweet) error
	GetTweetByID(id uint) (*tweetDomainModel.Tweet, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// SaveTweet saves a tweet in the database.
func (r *repository) SaveTweet(tweet *tweetDomainModel.Tweet) error {
	dbTweet := DBModel{
		Content:  tweet.Content,
		UserID:   tweet.UserID,
		Username: tweet.Username,
	}

	return r.db.Create(&dbTweet).Error
}

// GetTweetByID returns a tweet from the database by its ID, if the tweet does not exist it returns nil.
// If there is an error it returns the error.
func (r *repository) GetTweetByID(id uint) (*tweetDomainModel.Tweet, error) {
	var tweet DBModel

	err := r.db.Where("id = ?", id).First(&tweet).Error
	if err != nil {
		return nil, err
	}

	return tweet.toDomainModel(), nil
}
