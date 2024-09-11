package handler

import (
	"context"
	"smart-waste/apis/user/models/req"
	"smart-waste/domain/user/entity"
	"smart-waste/pkgs/res"
	"smart-waste/pkgs/security"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateUser Handles creating
func (u UserHandler) HandlerCreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		var createUserReq = new(req.CreateUserReq)
		if err := c.BodyParser(&createUserReq); err != nil {
			res := res.NewRes(
				fiber.StatusBadRequest,
				"Failed to parse request body. Please check the format of your input data.",
				false,
				nil,
			)
			res.SetError(err)
			return res.Send(c)
		}

		var _, err = u.GetUserByPhoneUsecase.ExecuteGetUserByPhone(ctx, createUserReq.Phone)
		if err != nil {
			// User already exists
			res := res.NewRes(fiber.StatusConflict, "User with the same phone already exists", false, nil)
			res.SetError(err)
			return res.Send(c)
		}

		hashedPassword, err := security.HashAndSalt([]byte(*createUserReq.Password))
		if err != nil {
			res := res.NewRes(fiber.StatusInternalServerError, "Failed to hash and salt password", false, nil)
			res.SetError(err)
			return res.Send(c)
		}

		userID, err := uuid.NewV7()
		if err != nil {
			res := res.NewRes(fiber.StatusInternalServerError, "Failed to generate UUID", false, nil)
			res.SetError(err)
			return res.Send(c)
		}

		var userEntity = entity.User{
			ID:        userID.String(),
			FirstName: createUserReq.FirstName,
			LastName:  createUserReq.LastName,
			Gender:    createUserReq.Gender,
			Role:      createUserReq.Role,
			Category:  createUserReq.Category,
			Email:     createUserReq.Email,
			Phone:     createUserReq.Phone,
			Password:  &hashedPassword,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		var useCaseErr = u.CreateUserUsecase.ExecuteCreateUser(ctx, &userEntity)
		if useCaseErr != nil {
			res := res.NewRes(fiber.StatusInternalServerError, "Failed to create user", false, nil)
			res.SetError(useCaseErr)
			return res.Send(c)
		}

		res := res.NewRes(fiber.StatusOK, "User created successfully", true, userEntity)
		return res.Send(c)
	}
}
