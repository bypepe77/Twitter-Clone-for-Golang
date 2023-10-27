package repositories

import (
	"errors"

	"github.com/bypepe77/Twitter-Clone-for-Golang/internal/domain/user"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *user.User) error
	GetUserByUsername(username string) (*user.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// CreateUser creates a new user in the database.
func (r *userRepository) CreateUser(user *user.User) error {
	dbUser := UserDBModel{
		Username: user.Username,
		Password: user.Password,
	}

	return r.db.Create(&dbUser).Error
}

// GetUserByUsername returns a user from the database by their usernme, if the user does not exist it returns nil.
// If there is an error it returns the error.
func (r *userRepository) GetUserByUsername(username string) (*user.User, error) {
	var user UserDBModel

	err := r.db.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user.toDomainModel(), nil
}
