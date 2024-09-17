package user

import "time"

type UserDTO struct {
	StudentID   string    `json:"student_id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Mail        string    `json:"mail"`
	ReviewAvg   float32    `json:"review_avg"`
	ReviewCount int				`json:"review_count"`
	JoinedAt    time.Time `json:"joined_at"`
}

type CreateUserDTO struct {
	StudentID   string    `json:"student_id"`
	Name        string    `json:"name"`
	Mail        string    `json:"mail"`
}
