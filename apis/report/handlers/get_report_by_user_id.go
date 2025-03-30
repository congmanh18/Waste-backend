package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *ReportHandler) HandlerGetReportsByUserID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		userID := c.Params("user_id")
		if userID == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Missing user ID")
		}

		reports, err := h.GetReportsByUserIDUsecase.Execute(ctx, &userID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve reports by user")
		}

		// var reportIDs []string
		// for _, report := range *reports {
		// 	reportIDs = append(reportIDs, report.ID)
		// }

		return c.Status(fiber.StatusOK).JSON(reports)
	}
}
