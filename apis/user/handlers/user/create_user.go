package handler

import (
	"smart-waste/apis/user/models/req"
	"smart-waste/domain/user/entity"
	usecase "smart-waste/domain/user/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	CreateUserUsecase *usecase.CreateUserUsecase
}

func (u UserHandler) HandlerCreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var createUserReq = new(req.CreateUserReq)
		if err := c.BodyParser(&createUserReq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		var userEntity = entity.User{
			ID:        createUserReq.ID,
			FirstName: createUserReq.FirstName,
			LastName:  createUserReq.LastName,
			Gender:    createUserReq.Gender,
			Role:      createUserReq.Role,
			Category:  createUserReq.Category,
			Email:     createUserReq.Email,
			Phone:     createUserReq.Phone,
			Password:  createUserReq.Password,
		}

		var useCaseErr = u.CreateUserUsecase.ExecuteCreateUser(c.Context(), userEntity)
		if useCaseErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": useCaseErr.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "OK",
		})
	}
}
