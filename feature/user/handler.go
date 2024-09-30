package user

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sds-2/model"
)

type UserHandler struct {
	UserRepository UserRepository
}

func NewUserHandler(UserRepository UserRepository) *UserHandler {
	return &UserHandler{UserRepository}
}

func (h UserHandler) GetUserById(c *fiber.Ctx) error {
	userIdStr := c.Params("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	user, err := h.UserRepository.GetUserById(userId)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	return c.Status(200).JSON(user)
}

func (h UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.UserRepository.GetUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to get users",
		})
	}
	return c.Status(200).JSON(users)
}

func (h UserHandler) UpdateUser(c *fiber.Ctx) error {

	params := c.Params("id")
	userId, err := strconv.Atoi(params)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid user ID",
		})
	}

	var user model.User
	user, err = h.UserRepository.GetUserById(userId)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	user.ID = userId

	err = h.UserRepository.UpdateUser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to update user",
		})
	}

	return c.Status(204).JSON(fiber.Map{})
}
func (h UserHandler) DeleteUser(c *fiber.Ctx) error {
	userIdStr := c.Params("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	err = h.UserRepository.DeleteUser(userId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete user",
		})
	}
	return c.Status(204).JSON(fiber.Map{})
}

func (h UserHandler) GetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	user, err := h.UserRepository.GetUserByEmail(email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	if user == nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	return c.Status(200).JSON(user)
}
