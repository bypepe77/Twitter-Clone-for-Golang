package repositories

import (
	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/domain/user"
	"gorm.io/gorm"
)

type UserDBModel struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (UserDBModel) TableName() string {
	return "users"
}

func (u *UserDBModel) toUserDomainModel() *user.User {
	return user.NewUser(u.Username, u.Password, u.ID)
}
