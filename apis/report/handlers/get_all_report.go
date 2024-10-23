package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *ReportHandler) HandlerGetAllReports() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		reports, err := h.GetAllReportsUsecase.Execute(ctx)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve reports")
		}

		return c.Status(fiber.StatusOK).JSON(reports)
	}
}
