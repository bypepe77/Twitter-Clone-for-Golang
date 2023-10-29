package tweet

import (
	"time"

	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/domain/user"
)

type Tweet struct {
	ID            uint
	Content       string
	UserID        uint
	Username      string
	User          user.User
	LikesCount    int
	RetweetsCount int
	RepliesCount  int
	CreatedAt     time.Time
}
