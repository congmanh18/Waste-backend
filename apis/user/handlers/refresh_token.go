package handler

import (
	"smart-waste/pkgs/auth"
	"smart-waste/pkgs/res"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (u UserHandler) RefreshTokenHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		refreshToken := c.Get("Refresh-Token")
		if refreshToken == "" {
			res := res.NewRes(
				fiber.StatusUnauthorized,
				"Missing or invalid refresh token",
				false,
				nil,
			)
			return res.Send(c)
		}

		claims, err := auth.ParseToken(refreshToken, "secret-key")
		if err != nil {
			res := res.NewRes(
				fiber.StatusUnauthorized,
				"Invalid or expired refresh token",
				false,
				nil,
			)
			res.SetError(err)
			return res.Send(c)
		}

		// Generate new access token
		accessToken, err := auth.GenerateJWTToken(claims.UserId, claims.Role, time.Minute*15)
		if err != nil {
			res := res.NewRes(
				fiber.StatusInternalServerError,
				"Failed to generate access token",
				false,
				nil,
			)
			res.SetError(err)
			return res.Send(c)
		}

		res := res.NewRes(fiber.StatusOK, "Token refreshed successfully", true, accessToken)
		return res.Send(c)
	}
}
