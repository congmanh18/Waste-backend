package handler

import (
	"context"
	"smart-waste/pkgs/res"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *ReportHandler) HandlerGetAllReports1() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		reports, err := h.GetAllReportsUsecase.Execute(ctx)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve reports")
		}

		// Trả về danh sách các ID
		res := res.NewRes(fiber.StatusOK, "All", true, reports)
		return res.Send(c)
	}
}
