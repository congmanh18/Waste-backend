package handler

import (
	"context"
	"smart-waste/apis/report/model"
	"smart-waste/domain/report/entity"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	// Add a logging library if needed
	// "github.com/sirupsen/logrus"
)

func (h *ReportHandler) HandlerCreateReport() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Create a context with a timeout to prevent long-running operations.
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		// Parse the request payload into the ReportReq struct.
		var reportReq model.ReportReq
		if err := c.BodyParser(&reportReq); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
		}

		// Validate essential fields.
		if reportReq.UserID == nil || reportReq.WasteBinID == nil {
			return fiber.NewError(fiber.StatusBadRequest, "UserID and WasteBinID are required")
		}

		// Generate a new UUID for the report.
		reportID := uuid.New()

		// Create a new report entity.
		report := entity.Report{
			ID:          reportID.String(),
			UserID:      reportReq.UserID,
			WasteBinID:  reportReq.WasteBinID,
			Image:       reportReq.Image,
			Description: reportReq.Description,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		// Attempt to create the report using the use case.
		if err := h.CreateReportUsecase.ExecuteCreateReport(ctx, report); err != nil {
			// Log the error if logging is available.
			// h.logger.Errorf("Failed to create report: %v", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to create report")
		}

		// Return the created report as a JSON response.
		return c.Status(fiber.StatusOK).JSON(report)
	}
}
