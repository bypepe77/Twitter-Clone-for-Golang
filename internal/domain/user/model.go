package user

type User struct {
	ID       uint
	Username string
	Password string
	Token    string
}

func New(username, password, token string, userID uint) *User {
	return &User{
		ID:       userID,
		Username: username,
		Password: password,
		Token:    token,
	}
}

func CreatBasicUser(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}
