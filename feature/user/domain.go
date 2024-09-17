package user

import (
	"github.com/sds-2/model"
)

// UserRepository defines the interface for user data operations.
type UserRepository interface {
	FindByID(studentID string) (*model.User, error)
	FindAll() ([]model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(studentID string) error
}

// UserDomain defines the interface for user business operations.
type UserDomain interface {
	GetUserByID(studentID string) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(studentID string) error
}

// userDomain implements UserDomain interface.
type userDomain struct {
	repo UserRepository
}

// NewUserDomain creates a new UserDomain instance.
func NewUserDomain(repo UserRepository) UserDomain {
	return &userDomain{repo: repo}
}

// GetUserByID retrieves a user by ID.
func (s *userDomain) GetUserByID(studentID string) (*model.User, error) {
	return s.repo.FindByID(studentID)
}

// GetAllUsers retrieves all users.
func (s *userDomain) GetAllUsers() ([]model.User, error) {
	return s.repo.FindAll()
}

// CreateUser creates a new user.
func (s *userDomain) CreateUser(user *model.User) error {
	return s.repo.Create(user)
}

// UpdateUser updates an existing user.
func (s *userDomain) UpdateUser(user *model.User) error {
	return s.repo.Update(user)
}

// DeleteUser deletes a user by ID.
func (s *userDomain) DeleteUser(studentID string) error {
	return s.repo.Delete(studentID)
}
