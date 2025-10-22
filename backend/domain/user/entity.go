package user

import "time"

// User represents the core user entity in the domain
type User struct {
	ID        string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser creates a new User entity
func NewUser(id, email string) *User {
	now := time.Now()
	return &User{
		ID:        id,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

