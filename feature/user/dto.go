package user

import "time"

// UserDTO represents the user data transfer object.
type UserDTO struct {
	StudentID   string    `json:"student_id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Mail        string    `json:"mail"`
	ReviewID    string    `json:"review_id"` // Foreign key field
	JoinedAt    time.Time `json:"joined_at"`
}

type CreateUserDTO struct {
	StudentID   string    `json:"student_id"`
	Name        string    `json:"name"`
	Mail        string    `json:"mail"`
	ReviewID    string    `json:"review_id"`
}
