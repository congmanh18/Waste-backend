package handler

import (
	req "smart-waste/apis/wastebin/model"
	"smart-waste/domain/wastebin/entity"

	"github.com/gofiber/fiber/v2"
)

func (w WasteBinHandler) HandlerUpdateWasteBin() fiber.Handler {
	return func(c *fiber.Ctx) error {
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

		var useCaseErr = w.UpdateWasteBinUsecase.ExecuteUpdateWasteBin(c.Context(), updateWasteBinReq.ID, &wasteBinEntity)
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
