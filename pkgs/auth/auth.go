package auth

import (
	"smart-waste/pkgs/res"

	"github.com/gofiber/fiber/v2"
)

// Dời sang package middleware
func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(res.NewRes(
			fiber.StatusUnauthorized,
			"Missing or invalid token",
			false,
			nil,
		))
	}

	claims, err := ParseToken(tokenString, "your-secret-key")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(res.NewRes(
			fiber.StatusUnauthorized,
			"Invalid or expired token",
			false,
			nil,
		))
	}

	// Attach user claims to the context
	c.Locals("id", claims.UserId)
	c.Locals("role", claims.Role)

	return c.Next()
}
