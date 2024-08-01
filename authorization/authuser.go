package authorization

// User contains user's information
type User struct {
	Username       string
	Email          string
	GoogleID       string
	GoogleEmail    string
	GoogleVerified bool
}

// NewUserAuth returns a new user with Google OAuth2 information
func NewUserAuth(username string, email string, googleID string, googleEmail string, googleVerified bool) (*User, error) {
	user := &User{
		Username:       username,
		Email:          email,
		GoogleID:       googleID,
		GoogleEmail:    googleEmail,
		GoogleVerified: googleVerified,
	}

	return user, nil
}