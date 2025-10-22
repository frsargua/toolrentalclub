package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/yourusername/toolrentalclub/domain/user"
)

// UserRepository implements user.Repository interface using in-memory storage
type UserRepository struct {
	mu    sync.RWMutex
	users map[string]*user.User // key is user ID
	index map[string]string     // email -> user ID index
}

// NewUserRepository creates a new in-memory user repository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*user.User),
		index: make(map[string]string),
	}
}

// FindByID retrieves a user by their ID
func (r *UserRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

// FindByEmail retrieves a user by their email
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	userID, exists := r.index[email]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	return r.users[userID], nil
}

// Create creates a new user
func (r *UserRepository) Create(ctx context.Context, user *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if user already exists
	if _, exists := r.users[user.ID]; exists {
		return fmt.Errorf("user already exists")
	}

	// Check if email is already taken
	if _, exists := r.index[user.Email]; exists {
		return fmt.Errorf("email already taken")
	}

	r.users[user.ID] = user
	r.index[user.Email] = user.ID

	return nil
}

// Update updates an existing user
func (r *UserRepository) Update(ctx context.Context, user *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if user exists
	existingUser, exists := r.users[user.ID]
	if !exists {
		return fmt.Errorf("user not found")
	}

	// If email changed, update index
	if existingUser.Email != user.Email {
		delete(r.index, existingUser.Email)
		r.index[user.Email] = user.ID
	}

	r.users[user.ID] = user

	return nil
}

