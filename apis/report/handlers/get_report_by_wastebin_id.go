package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)


func (h *ReportHandler) HandlerGetReportsByWasteBinID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		wasteBinID := c.Params("wastebin_id")
		if wasteBinID == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Missing wastebin ID")
		}

		reports, err := h.GetReportsByWasteBinIDUsecase.Execute(ctx, &wasteBinID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve reports by wastebin")
		}

		return c.Status(fiber.StatusOK).JSON(reports)
	}
}
