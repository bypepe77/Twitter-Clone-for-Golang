package user

//nolint: deadcode,unused
type Repository interface {
	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)
}
