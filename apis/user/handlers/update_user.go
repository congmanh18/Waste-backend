package handler

import (
	"context"
	"smart-waste/apis/user/models/req"
	"smart-waste/domain/user/entity"
	"smart-waste/pkgs/res"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (u UserHandler) HandlerUpdateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		var updateUserReq = new(req.UpdateUserReq)
		if err := c.BodyParser(&updateUserReq); err != nil {
			res := res.NewRes(
				fiber.StatusBadRequest,
				"Failed to parse request body. Please check the format of your input data.",
				false,
				nil,
			)
			res.SetError(err)
			return res.Send(c)
		}

		var userEntity = entity.User{
			FirstName: updateUserReq.FirstName,
			LastName:  updateUserReq.LastName,
			Gender:    updateUserReq.Gender,
			Category:  updateUserReq.Category,
			Email:     updateUserReq.Email,
			Password:  updateUserReq.Password,
		}

		var useCaseErr = u.UpdateUserUsecase.ExecuteUpdateUser(ctx, updateUserReq.ID, &userEntity)
		if useCaseErr != nil {
			res := res.NewRes(fiber.StatusInternalServerError, useCaseErr.Error(), false, nil)
			return res.Send(c)
		}

		res := res.NewRes(fiber.StatusOK, "User updated successfully", true, nil)
		return res.Send(c)
	}
}
