package handler

import (
	"context"
	"smart-waste/apis/report/model"
	"smart-waste/domain/report/entity"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *ReportHandler) HandlerCreateReport() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		var reportReq model.ReportReq
		if err := c.BodyParser(&reportReq); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
		}

		reportID := uuid.New().String()
		report := entity.Report{
			ID:          reportID,
			UserID:      reportReq.UserID,
			WasteBinID:  reportReq.WasteBinID,
			Image:       reportReq.Image,
			Description: reportReq.Description,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if err := h.CreateReportUsecase.ExecuteCreateReport(ctx, report); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to create report")
		}

		return c.Status(fiber.StatusOK).JSON(report)
	}
}
