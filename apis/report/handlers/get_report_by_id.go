package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *ReportHandler) HandlerGetReportByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		reportID := c.Params("id")
		if reportID == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Missing report ID")
		}

		report, err := h.GetReportByIDUsecase.Execute(ctx, &reportID)
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "Report not found")
		}

		return c.Status(fiber.StatusOK).JSON(report)
	}
}
