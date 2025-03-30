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

		// Lấy ID từ URL params
		id := c.Params("id")
		if id == "" {
			res := res.NewRes(
				fiber.StatusBadRequest,
				"Missing user ID in URL parameters.",
				false,
				nil,
			)
			return res.Send(c)
		}

		// Parse request body
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

		// Tạo entity từ request body
		var userEntity = entity.User{
			FirstName: updateUserReq.FirstName,
			LastName:  updateUserReq.LastName,
			Gender:    updateUserReq.Gender,
			Category:  updateUserReq.Category,
			Email:     updateUserReq.Email,
			Password:  updateUserReq.Password,
		}

		// Gọi usecase để update user
		var useCaseErr = u.UpdateUserUsecase.ExecuteUpdateUser(ctx, id, &userEntity)
		if useCaseErr != nil {
			res := res.NewRes(fiber.StatusInternalServerError, useCaseErr.Error(), false, nil)
			return res.Send(c)
		}

		// Trả về response thành công
		res := res.NewRes(fiber.StatusOK, "User updated successfully", true, nil)
		return res.Send(c)
	}
}
