package auth

const (
	// ErrInvalidCredentials is returned when the credentials are invalid
	ErrInvalidCredentials = "invalid credentials"
	// ErrUserNotFound is returned when the user is not found
	ErrUserNotFound = "user not found"
	// ErrUserAlreadyExists is returned when the user already exists
	ErrUserAlreadyExists = "user already exists"
	// ErrInvalidPayload is returned when the payload is invalid
	ErrInvalidPayload = "invalid payload"
	// ErrEmptyPassword is returned when the username is empty
	ErrEmptyUsername = "empty username"
	// ErrEmptyPassword is returned when the password is empty
	ErrEmptyPassword = "empty password"
	// ErrInternalServer is returned when the server has an internal error
	ErrInternalServer = "internal server error"
)

type input struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type response struct {
	Token    string `json:"token"`
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func newUserResponse(token, username string, userID uint) *response {
	return &response{
		Token:    token,
		ID:       userID,
		Username: username,
	}
}
