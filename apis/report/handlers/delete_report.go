package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *ReportHandler) HandlerDeleteReport() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		reportID := c.Params("id")
		if reportID == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Missing report ID")
		}

		if err := h.DeleteReportUsecase.Execute(ctx, &reportID); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete report")
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Report deleted successfully",
		})
	}
}
