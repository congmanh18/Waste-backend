package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *ReportHandler) HandlerGetReportsByDate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		date := c.Query("date")
		if date == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Missing date parameter")
		}

		reports, err := h.GetReportsByDateUsecase.Execute(ctx, &date)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve reports by date")
		}

		return c.Status(fiber.StatusOK).JSON(reports)
	}
}
