package handler

import (
	"context"
	"fmt"
	"smart-waste/apis/user/models/req"
	tokenRes "smart-waste/apis/user/models/res"
	"smart-waste/pkgs/auth"
	"smart-waste/pkgs/res"
	"smart-waste/pkgs/security"
	validate "smart-waste/pkgs/validator"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func (u UserHandler) HandlerLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		// Parse request body to get login user request
		var loginUserReq = new(req.LoginUserReq)
		if err := c.BodyParser(&loginUserReq); err != nil {
			res := res.NewRes(
				fiber.StatusBadRequest,
				"Failed to parse login user request",
				false,
				nil,
			)
			res.SetError(err)
			return res.Send(c)
		}

		// Validate các field trong LoginUserReq
		err := validate.Validate.Struct(loginUserReq)
		if err != nil {
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				var errorMessages []string
				for _, validationErr := range validationErrors {
					msg := fmt.Sprintf("Field: %s, Error: %s", validationErr.Field(), validationErr.Tag())
					errorMessages = append(errorMessages, msg)
				}

				// Trả về danh sách các lỗi validation...
				// Note chú ý data là lỗi
				res := res.NewRes(
					fiber.StatusBadRequest,
					"Validation failed",
					false,
					errorMessages,
				)
				return res.Send(c)
			}

			res := res.NewRes(
				fiber.StatusInternalServerError,
				"Internal server error",
				false,
				nil,
			)

			res.SetError(err)
			return res.Send(c)
		}

		// Find user by phone number
		var foundUser, useCaseErr = u.GetUserByPhoneUsecase.ExecuteGetUserByPhone(ctx, &loginUserReq.Phone)
		if useCaseErr != nil {
			res := res.NewRes(fiber.StatusInternalServerError, "Error while finding user", false, nil)
			res.SetError(useCaseErr)
			return res.Send(c)
		}

		if foundUser == nil {
			res := res.NewRes(fiber.StatusUnauthorized, "User not found", false, nil)
			return res.Send(c)
		}

		// Compare password
		if !security.ComparePasswords(*foundUser.Password, []byte(*loginUserReq.Password)) {
			res := res.NewRes(
				fiber.StatusUnauthorized,
				"Invalid credentials",
				false,
				nil,
			)
			return res.Send(c)
		}

		// Create JWT custom claims
		claims := auth.JwtCustomClaims{
			ID:    foundUser.ID,
			Role:  *foundUser.Role,
			Phone: *foundUser.Phone,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 15).Unix(), // 15 mins expiration
			},
		}

		// Create JWT custom refresh_token claims
		rf_claims := auth.JwtCustomClaims{
			ID: foundUser.ID,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // 15 mins expiration
			},
		}

		// Generate JWT tokens with claims
		accessToken, err := auth.GenerateTokenWithClaims(claims)
		if err != nil {
			res := res.NewRes(fiber.StatusInternalServerError, "Failed to generate access token", false, nil)
			res.SetError(err)
			return res.Send(c)
		}

		// Refresh token with longer expiration
		refreshToken, err := auth.GenerateTokenWithClaims(rf_claims)
		if err != nil {
			res := res.NewRes(fiber.StatusInternalServerError, "Failed to generate refresh token", false, nil)
			res.SetError(err)
			return res.Send(c)
		}

		// Store the JWT token and refresh token in the database
		foundUser.Token = &accessToken
		foundUser.RefreshToken = &refreshToken
		if err := u.UpdateUserUsecase.ExecuteUpdateUser(ctx, foundUser.ID, foundUser); err != nil {
			res := res.NewRes(fiber.StatusInternalServerError, "Failed to update user tokens in database", false, nil)
			res.SetError(err)
			return res.Send(c)
		}

		// Send response with both tokens
		tokenResponse := tokenRes.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		res := res.NewRes(fiber.StatusOK, "User login successful", true, tokenResponse)
		return res.Send(c)
	}
}
