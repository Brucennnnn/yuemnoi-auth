package user

import (
	"fmt"

	"github.com/sds-2/model"
)

type UserDomain interface {
	GetUserByID(studentID string) (model.User, error)
	GetAllUsers() ([]model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(studentID string) error
}

type userDomain struct {
	repo UserRepository
}

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

func (s *userDomain) DeleteUser(studentID string) error {
	err := s.repo.DeleteUser(studentID)
	if err != nil {
		return fmt.Errorf("failed to delete user from user domain: %w", err)
	}
	return nil
}
