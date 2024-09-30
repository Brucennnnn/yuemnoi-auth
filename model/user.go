package model

type User struct {
	ID      int      `json:"id" gorm:"primaryKey"`
	Email   string   `json:"email" gorm:"<-:create;unique;not null"`
	Name    string   `json:"name"`
	SurName string   `json:"surname"`
    ReviewsGiven  []Review `json:"reviews_given" gorm:"foreignKey:ReviewerID"`
    ReviewsReceived []Review `json:"reviews_received" gorm:"foreignKey:RevieweeID"`
}
