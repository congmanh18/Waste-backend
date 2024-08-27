package handler

import "github.com/gofiber/fiber/v2"

func (u UserHandler) HandlerFindUserByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userEntity, err := u.FindUserByIDUsecase.ExecuteFindById(c.Context(), c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(userEntity)
	}
}
