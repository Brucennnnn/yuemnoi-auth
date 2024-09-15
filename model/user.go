package model

type User struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	SurName string `json:"surname"`
	Email   string `json:"email"`
}
