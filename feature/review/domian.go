package review

import (
	"fmt"

	"github.com/sds-2/model"
)

type ReviewDomain interface {
	GetReviewsByUserId(userId int) ([]model.Review, error)
	CreateReview(review model.Review) (model.Review, error)
}

type ReviewDomainImpl struct {
	reviewRepository ReviewRepository
}

func NewReviewDomain(reviewRepository ReviewRepository) *ReviewDomainImpl {
	return &ReviewDomainImpl{reviewRepository: reviewRepository}
}

func (i *ReviewDomainImpl) GetReviewsByUserId(userId int) ([]model.Review, error) {
	fmt.Println("hell")
	reviews, err := i.reviewRepository.GetReviewsByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get reviews from review domain: %w", err)
	}
	return reviews, nil
}

func (i *ReviewDomainImpl) CreateReview(review model.Review) (model.Review, error) {
	review, err := i.reviewRepository.CreateReview(review)
	if err != nil {
		return model.Review{}, fmt.Errorf("failed to create review from review domain: %w", err)
	}
	return review, nil
}
