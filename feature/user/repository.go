package user

import (
	"github.com/sds-2/model"
	"gorm.io/gorm"
)

// userRepository implements UserRepository interface.
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new userRepository instance.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// FindByID retrieves a user by ID.
func (r *userRepository) FindByID(studentID string) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, "student_id = ?", studentID).Error
	return &user, err
}

// FindAll retrieves all users.
func (r *userRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

// Create adds a new user to the database.
func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// Update updates an existing user.
func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

// Delete removes a user by ID.
func (r *userRepository) Delete(studentID string) error {
	return r.db.Delete(&model.User{}, "student_id = ?", studentID).Error
}
