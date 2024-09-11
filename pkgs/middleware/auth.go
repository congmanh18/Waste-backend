package middleware

import (
	"smart-waste/pkgs/auth"
	"smart-waste/pkgs/res"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Get the token from Authorization header
	tokenString := c.Get("Authorization")
	if tokenString == "" || len(tokenString) < 7 || tokenString[:7] != "Bearer " {
		return c.Status(fiber.StatusUnauthorized).JSON(res.NewRes(
			fiber.StatusUnauthorized,
			"Missing or invalid token",
			false,
			nil,
		))
	}

	// Remove "Bearer " prefix from token string
	tokenString = tokenString[7:]

	// Parse and validate token
	claims, err := auth.ParseToken(tokenString, "your-secret-key")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(res.NewRes(
			fiber.StatusUnauthorized,
			"Invalid or expired token",
			false,
			nil,
		))
	}

	// Attach user claims to the context (id, role)
	c.Locals("id", claims.ID)
	c.Locals("role", claims.Role)

	return c.Next() // Continue to next handler
}
