package middlewares

import (
	"strings"

	"github.com/irfanalmsyah/ecommerce-api/helpers"

	"github.com/gofiber/fiber/v2"
)

func Protect(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return helpers.UnauthorizedError(c, "Missing authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return helpers.UnauthorizedError(c, "Invalid authorization header format")
	}

	tokenStr := parts[1]

	claims, err := helpers.VerifyToken(tokenStr)
	if err != nil {
		return helpers.UnauthorizedError(c, err.Error())
	}

	userID := uint((*claims)["user_id"].(float64))

	c.Locals("userID", userID)

	return c.Next()
}
