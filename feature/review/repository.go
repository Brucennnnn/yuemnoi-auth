package review

import (
	"github.com/sds-2/model"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetReviewsByUserId(userId int) ([]model.Review, error)
	CreateReview(review model.Review) (error)
}

type ReviewRepositoryImpl struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepositoryImpl {
	return &ReviewRepositoryImpl{
		db: db,
	}
}
func (i ReviewRepositoryImpl) GetReviewsByUserId(userId int) ([]model.Review, error) {
	var reviews []model.Review
	if err := i.db.Where("reviewee_id = ?", userId).Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (i ReviewRepositoryImpl) CreateReview(review model.Review) (error) {
	if err := i.db.Create(&review).Error; err != nil {
		return err
	}
	return  nil
}
