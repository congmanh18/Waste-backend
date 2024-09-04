package handler

import "github.com/gofiber/fiber/v2"

func (w WasteBinHandler) HandlerDeleteWasteBin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := w.DeleteWasteBinUsecase.ExecuteDeleteWasteBin(c.Context(), c.Params("id"))

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "Waste Bin deleted successfully",
		})
	}
}
