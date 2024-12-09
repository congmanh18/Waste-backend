package handler

import (
	"context"
	"smart-waste/pkgs/res"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *ReportHandler) HandlerGetLast() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		reportID := c.Params("id")
		if reportID == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Missing report ID")
		}

		report, err := h.GetLast.Execute(ctx, &reportID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve reports")
		}

		// Trả về danh sách các ID
		res := res.NewRes(fiber.StatusOK, "All", true, report)
		return res.Send(c)
	}
}
