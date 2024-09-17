package user

import (
	"github.com/sds-2/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByID(studentID string) (model.User, error)
	FindAllUsers() ([]model.User, error)
	CreateUser(user model.User) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
	DeleteUser(studentID string) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindUserByID(studentID string) (model.User, error) {
	var user model.User
	err := r.db.First(&user, "student_id = ?", studentID).Error
	return &user, err
}

func (r *userRepository) FindAllUsers() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user model.User) (model.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) DeleteUser(studentID string) error {
	if err := r.db.Delete(&model.User{}, "student_id = ?", studentID).Error; err != nil {
		return err
	}
	return nil
}
