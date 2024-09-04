package handler

import "github.com/gofiber/fiber/v2"

func (w WasteBinHandler) HandlerReadWasteBin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		wateBinEntity, err := w.ReadWasteBinUsecase.ReadWasteBinByID(c.Context(), c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(wateBinEntity)
	}
}
