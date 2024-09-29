package user

import (
	"fmt"

	"github.com/sds-2/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserById(userId int) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUsers() ([]model.User, error)
	UpdateUser(user model.User) error
	DeleteUser(userId int) error
	GetUserByEmail(email string) (*model.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (i UserRepositoryImpl) GetUserById(userId int) (model.User, error) {
	var user model.User
	if err := i.db.Where("id = ?", userId).Find(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (i UserRepositoryImpl) CreateUser(user model.User) (model.User, error) {
	if err := i.db.Create(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (i UserRepositoryImpl) UpdateUser(userInfo model.User) error {

	if err := i.db.Save(&userInfo).Error; err != nil {
		return fmt.Errorf("failed to update user")
	}
	return nil
}

func (i UserRepositoryImpl) DeleteUser(userId int) error {
	if err := i.db.Where("id = ?", userId).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (i UserRepositoryImpl) GetUsers() ([]model.User, error) {
	var users []model.User
	if err := i.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (i UserRepositoryImpl) GetUserByEmail(email string) (*model.User, error) {
	var user *model.User
	result := i.db.Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return user, nil
}
