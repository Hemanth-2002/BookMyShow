package authorization

// User contains user's information
type User struct {
	Username string
	Email    string
}

// NewUser returns a new user
func NewUserAuth(username string, email string) (*User, error) {
	user := &User{
		Username: username,
		Email:    email,
	}

	return user, nil
}
