package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sds-2/model"
)

// UserHandler handles HTTP requests related to users.
type UserHandler struct {
	service UserDomain
}

// NewUserHandler creates a new UserHandler instance.
func NewUserHandler(service UserDomain) *UserHandler {
	return &UserHandler{service: service}
}

// CreateUser handles the creation of a new user.
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var userDTO CreateUserDTO
	if err := c.BodyParser(&userDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	var reviewID *string
	if userDTO.ReviewID != "" {
			reviewID = &userDTO.ReviewID
	}

	user := model.User{
		StudentID:   userDTO.StudentID,
		Name:        userDTO.Name,
		Mail:        userDTO.Mail,
		ReviewID:		 reviewID,
	}

	if err := h.service.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(userDTO)
}

// GetUsers handles fetching all users.
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	var userDTOs []UserDTO
    for _, user := range users {
        var reviewID string
        if user.ReviewID != nil {
            reviewID = *user.ReviewID
        }

        userDTOs = append(userDTOs, UserDTO{
            StudentID:   user.StudentID,
            Name:        user.Name,
            PhoneNumber: user.PhoneNumber,
            Mail:        user.Mail,
            ReviewID:    reviewID,
            JoinedAt:    user.JoinedAt,
        })
    }

	return c.JSON(userDTOs)
}

// GetUser handles fetching a user by ID.
func (h *UserHandler) ViewUserProfile(c *fiber.Ctx) error {
	studentID := c.Params("student_id")
	user, err := h.service.GetUserByID(studentID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	var reviewID string
    if user.ReviewID != nil {
        reviewID = *user.ReviewID
    }

    userDTO := UserDTO{
        StudentID:   user.StudentID,
        Name:        user.Name,
        PhoneNumber: user.PhoneNumber,
        Mail:        user.Mail,
        ReviewID:    reviewID, // Convert *string to string
        JoinedAt:    user.JoinedAt,
    }

	return c.JSON(userDTO)
}

// UpdateUser handles updating an existing user.
func (h *UserHandler) ManageUserProfile(c *fiber.Ctx) error {
	studentID := c.Params("student_id")
	var userDTO UserDTO
	if err := c.BodyParser(&userDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Fetch existing user data first
	user, err := h.service.GetUserByID(studentID)
	if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Update only the fields that are provided in the request
	if userDTO.PhoneNumber != "" {
			user.PhoneNumber = userDTO.PhoneNumber
	}

	if userDTO.ReviewID != "" {
			reviewID := userDTO.ReviewID
			user.ReviewID = &reviewID // Convert string to *string
	}

	// Update the user in the database
	if err := h.service.UpdateUser(studentID, user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.JSON(fiber.Map{
			"message": "User updated successfully",
      "studentID": studentID,
	})
}


// DeleteUser handles deleting a user by ID.
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	studentID := c.Params("student_id")
	if err := h.service.DeleteUser(studentID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
