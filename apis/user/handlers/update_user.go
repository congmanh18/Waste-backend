package handler

import (
	"smart-waste/apis/user/models/req"
	"smart-waste/domain/user/entity"

	"github.com/gofiber/fiber/v2"
)

func (u UserHandler) HandlerUpdateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var updateUserReq = new(req.UpdateUserReq)
		if err := c.BodyParser(&updateUserReq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		var userEntity = entity.User{
			FirstName: updateUserReq.FirstName,
			LastName:  updateUserReq.LastName,
			Gender:    updateUserReq.Gender,
			Category:  updateUserReq.Category,
			Email:     updateUserReq.Email,
			Password:  updateUserReq.Password,
		}

		var useCaseErr = u.UpdateUserUsecase.ExecuteUpdateUser(c.Context(), updateUserReq.ID, &userEntity)
		if useCaseErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": useCaseErr.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "User updated successfully",
		})
	}
}
