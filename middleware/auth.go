package middleware

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sds-2/config"
	"github.com/sds-2/model"
)

func AuthMiddleware(config *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if config.Environment == "dev" {
			return c.Next()
		}

		tokenString := c.Cookies(config.Cookie.CookieNameAuth)

		token, err := jwt.ParseWithClaims(tokenString, &model.AuthTokenClaim{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			secretKey := []byte(config.Cookie.Secret)
			return secretKey, nil
		})

		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}

		if claims, ok := token.Claims.(*model.AuthTokenClaim); ok && token.Valid {
			c.Locals("user_id", strconv.Itoa(claims.UserID))

		} else {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}

		return c.Next()
	}
}
