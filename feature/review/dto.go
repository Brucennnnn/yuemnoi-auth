package review

type ReviewUserRequest struct {
	Rating      int `json:"rating"`
	Description string `json:"description"`
	ReviewerID  int    `json:"reviewer_id"` 
	RevieweeID  int    `json:"reviewee_id"`
}

type ViewUserReviewsResponse struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	Rating       int    `json:"rating"`
	Description  string `json:"description"`
	ReviewerID   int    `json:"reviewer_id"`
}