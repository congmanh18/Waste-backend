package handler

import (
	req "smart-waste/apis/wastebin/model"
	"smart-waste/domain/wastebin/entity"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (u WasteBinHandler) HandlerCreateWasteBin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var createWasteBinReq = new(req.CreateWasteBinReq)
		if err := c.BodyParser(&createWasteBinReq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}
		// Validate request data here (e.g., weight, filledLevel, airQuality, waterLevel, address, latitude, longitude)
		wasteBinID := uuid.New().String()
		var wasteBinEntity = entity.WasteBin{
			ID:          wasteBinID,
			Weight:      createWasteBinReq.Weight,
			FilledLevel: createWasteBinReq.FilledLevel,
			AirQuality:  createWasteBinReq.AirQuality,
			WaterLevel:  createWasteBinReq.WaterLevel,
			Address:     createWasteBinReq.Address,
			Latitude:    createWasteBinReq.Latitude,
			Longitude:   createWasteBinReq.Longitude,
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message":  "Waste bin created successfully",
			"wasteBin": wasteBinEntity,
		})

	}
}
