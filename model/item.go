package model

type Item struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
