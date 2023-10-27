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

func (u *UserDBModel) toDomainModel() *user.User {
	return &user.User{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
	}
}
