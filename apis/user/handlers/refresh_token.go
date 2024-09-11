package handler

import (
	tokenRes "smart-waste/apis/user/models/res"
	"smart-waste/pkgs/auth"
	"smart-waste/pkgs/res"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func (u UserHandler) RefreshTokenHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the refresh token from request body
		var refreshTokenReq struct {
			RefreshToken string `json:"refreshToken"`
		}

		if err := c.BodyParser(&refreshTokenReq); err != nil {
			res := res.NewRes(fiber.StatusBadRequest, "Invalid request", false, nil)
			return res.Send(c)
		}

		// Parse the token claims to validate refresh token
		token, err := jwt.ParseWithClaims(refreshTokenReq.RefreshToken, &auth.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return auth.JwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			res := res.NewRes(fiber.StatusUnauthorized, "Invalid refresh token", false, nil)
			return res.Send(c)
		}

		claims, ok := token.Claims.(*auth.JwtCustomClaims)
		if !ok || !token.Valid {
			res := res.NewRes(fiber.StatusUnauthorized, "Invalid token claims", false, nil)
			return res.Send(c)
		}

		// Generate a new access token
		newAccessToken, err := auth.GenerateTokenWithClaims(*claims)
		if err != nil {
			res := res.NewRes(fiber.StatusInternalServerError, "Failed to generate new access token", false, nil)
			return res.Send(c)
		}

		// Respond with the new access token
		tokenResponse := tokenRes.TokenResponse{
			AccessToken: newAccessToken,
		}

		res := res.NewRes(fiber.StatusOK, "Access token refreshed successfully", true, tokenResponse)
		return res.Send(c)
	}
}
