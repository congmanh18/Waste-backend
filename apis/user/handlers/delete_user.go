package handler

import "github.com/gofiber/fiber/v2"

func (u UserHandler) HandlerDeleteUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := u.DeleteUserUsecase.ExecuteDeleteUser(c.Context(), c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "User deleted successfully",
		})
	}
}
