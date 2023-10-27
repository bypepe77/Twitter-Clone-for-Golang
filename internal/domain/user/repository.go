package user

type Repository interface {
	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)
}
