package user

type User struct {
	ID       uint
	Username string
	Password string
}

func NewUser(username, password string, userID uint) *User {
	return &User{
		ID:       userID,
		Username: username,
		Password: password,
	}
}
