package review

import (
	"fmt"
	"github.com/sds-2/model"
)

type ReviewDomain interface {
	ViewUserReviews(userId int) ([]model.Review, error)
	ReviewUser(review model.Review) (model.Review, error)
}

type reviewDomainImpl struct {
	reviewRepository
}

func NewReviewDomain(reviewRepository reviewRepository) *reviewDomainImpl {
	return &reviewDomainImpl{reviewRepository: reviewRepository}
}

func (i reviewDomainImpl) ViewUserReviews(userId int) ([]model.Review, error) {
	reviews, err := i.reviewRepository.ViewUserReviews(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get reviews from review domain: %w", err)
	}
	return reviews, nil
}

func (i reviewDomainImpl) ReviewUser(review model.Review) ( model.Review, error) {
	review, err := i.reviewRepository.ReviewUser(review)
	if err != nil {
		return model.Review{}, fmt.Errorf("failed to create review from review domain: %w", err)
	}
	return review, nil
}

