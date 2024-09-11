package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (w WasteBinHandler) HandlerReadWasteBin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		wateBinEntity, err := w.ReadWasteBinUsecase.ReadWasteBinByID(ctx, c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(wateBinEntity)
	}
}
