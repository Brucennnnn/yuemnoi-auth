package user

import (
	"fmt"

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
	fmt.Println('1')

	if err := c.BodyParser(&userDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	user := model.User{
		StudentID: userDTO.StudentID,
		Name:      userDTO.Name,
		Mail:      userDTO.Mail,
	}

	return c.Status(fiber.StatusCreated).JSON(userDTO)
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers() // Ensure h.service is properly initialized
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	var userDTOs []UserDTO
	for _, user := range users {
		userDTOs = append(userDTOs, UserDTO{
			StudentID:   user.StudentID,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
			Mail:        user.Mail,
			JoinedAt:    user.JoinedAt,
		})
	}

	return c.JSON(userDTOs)
}

func (h *UserHandler) ViewUserProfile(c *fiber.Ctx) error {
	studentID := c.Params("student_id")
	user, err := h.service.GetUserByID(studentID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	userDTO := UserDTO{
		StudentID:   user.StudentID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Mail:        user.Mail,
		JoinedAt:    user.JoinedAt,
	}

	return c.JSON(userDTO)
}

func (h *UserHandler) ManageUserProfile(c *fiber.Ctx) error {
	studentID := c.Params("student_id")
	var userDTO UserDTO
	if err := c.BodyParser(&userDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	user, err := h.service.GetUserByID(studentID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Update the user in the database
	if err := h.service.UpdateUser(&model.User{
		StudentID:   studentID,
		Name:        userDTO.Name,
		PhoneNumber: userDTO.PhoneNumber,
		Mail:        userDTO.Mail,
		JoinedAt:    user.JoinedAt,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.JSON(fiber.Map{
		"message":   "User updated successfully",
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
