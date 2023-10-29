package tweet

import (
	tweetDomainModel "github.com/bypepe77/Twitter-Clone-for-Golang/internal/domain/tweet"
	user "github.com/bypepe77/Twitter-Clone-for-Golang/internal/infrastructure/repositories/user"
	"gorm.io/gorm"
)

type DBModel struct {
	gorm.Model
	Content       string           `gorm:"type:varchar(140);not null"`
	UserID        uint             `gorm:"not null"`
	Username      string           `gorm:"not null"`
	User          user.UserDBModel `gorm:"foreignkey:UserID"`
	LikesCount    int              `gorm:"not null;default:0"`
	RetweetsCount int              `gorm:"not null;default:0"`
	RepliesCount  int              `gorm:"not null;default:0"`
}

func (DBModel) TableName() string {
	return "tweets"
}

func (d *DBModel) toDomainModel() *tweetDomainModel.Tweet {
	return &tweetDomainModel.Tweet{
		ID:            d.ID,
		Content:       d.Content,
		UserID:        d.UserID,
		Username:      d.Username,
		LikesCount:    d.LikesCount,
		RetweetsCount: d.RetweetsCount,
		RepliesCount:  d.RepliesCount,
		CreatedAt:     d.CreatedAt,
	}
}
