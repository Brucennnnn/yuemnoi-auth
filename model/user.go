package model

import (
	"time"
)

type User struct {
	StudentID   string    `gorm:"primaryKey;not null;unique" json:"student_id"`
	Name        string    `gorm:"not null" json:"name"`
	PhoneNumber string    `gorm:"type:varchar(10)" json:"phone_number"`
	Mail        string    `gorm:"not null" json:"mail"`
	JoinedAt    time.Time `gorm:"type:timestamptz;not null;default:now()" json:"joined_at"`
}
