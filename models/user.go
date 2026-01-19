package models

import "time"

// User represents a system user
type User struct {
	ID        int
	Username  string
	Email     string
	Password  string // In real app: hashed password
	CreatedAt time.Time
}

// NewUser creates a new user instance
func NewUser(id int, username, email, password string) User {
	return User{
		ID:        id,
		Username:  username,
		Email:     email,
		Password:  password, // TODO: Hash in production
		CreatedAt: time.Now(),
	}
}

// Authenticate validates user credentials (dummy implementation)
func (u *User) Authenticate(password string) bool {
	// TODO: Implement proper password hashing
	return u.Password == password
}
