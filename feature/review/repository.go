package review

import (
	"github.com/sds-2/model"
	"gorm.io/gorm"
)

type reviewRepository interface {
	ViewUserReviews(userId int) ([]model.Review, error)
	ReviewUser(review model.Review) (model.Review, error)
}

type reviewRepositoryImpl struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *reviewRepositoryImpl {
	return &reviewRepositoryImpl{
		db: db,
	}
}
func (i reviewRepositoryImpl) ViewUserReviews(userId int) ([]model.Review, error) {
	var reviews []model.Review
	if err := i.db.Where("reviewee_id = ?", userId).Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}


func (i reviewRepositoryImpl) ReviewUser(review model.Review) (model.Review, error) {
	if err := i.db.Create(&review).Error; err != nil {
		return model.Review{}, err
	}
	return review, nil
}
