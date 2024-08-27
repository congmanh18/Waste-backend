package handler

import "github.com/gofiber/fiber/v2"

func (u UserHandler) HandlerFindAllUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userList, err := u.FindAllUserUsecase.ExecuteFindAll(c.Context())

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(userList)
	}
}
