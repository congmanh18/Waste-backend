package handler

import (
	"context"
	req "smart-waste/apis/wastebin/models"
	"smart-waste/domain/wastebin/entity"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (w WasteBinHandler) HandlerUpdateWasteBin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		var updateWasteBinReq = new(req.CreateWasteBinReq)
		if err := c.BodyParser(&updateWasteBinReq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		var wasteBinEntity = entity.WasteBin{
			Weight:      updateWasteBinReq.Weight,
			FilledLevel: updateWasteBinReq.FilledLevel,
			AirQuality:  updateWasteBinReq.AirQuality,
			WaterLevel:  updateWasteBinReq.WaterLevel,
			Address:     updateWasteBinReq.Address,
			Latitude:    updateWasteBinReq.Latitude,
			Longitude:   updateWasteBinReq.Longitude,
		}

		var useCaseErr = w.UpdateWasteBinUsecase.ExecuteUpdateWasteBin(ctx, updateWasteBinReq.ID, &wasteBinEntity)
		if useCaseErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"errors": useCaseErr.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "Waste bin updated successfully",
		})
	}
}
