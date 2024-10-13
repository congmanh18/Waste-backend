package handler

import (
	"context"
	req "smart-waste/apis/wastebin/models"
	"smart-waste/domain/wastebin/entity"
	"smart-waste/pkgs/res"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (u WasteBinHandler) HandlerCreateWasteBin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		var createWasteBinReq = new(req.CreateWasteBinReq)
		if err := c.BodyParser(&createWasteBinReq); err != nil {
			res := res.NewRes(
				fiber.StatusBadRequest,
				"Failed to parse request body. Please check the format of your input data.",
				false,
				nil,
			)
			res.SetError(err)
			return res.Send(c)
		}

		// Validate request data here (e.g., weight, filledLevel, airQuality, waterLevel, address, latitude, longitude)

		wasteBinID, _ := uuid.NewV7()

		var wasteBinEntity = entity.WasteBin{
			ID:            wasteBinID.String(),
			Weight:        createWasteBinReq.Weight,
			RemainingFill: createWasteBinReq.RemainingFill,
			AirQuality:    createWasteBinReq.AirQuality,
			WaterLevel:    createWasteBinReq.WaterLevel,
			Address:       createWasteBinReq.Address,
			Latitude:      createWasteBinReq.Latitude,
			Longitude:     createWasteBinReq.Longitude,
			Timestamp:     time.Now(),
		}

		var useCaseErr = u.CreateWasteBinUsecase.ExecuteCreateWasteBin(ctx, &wasteBinEntity)
		if useCaseErr != nil {
			res := res.NewRes(fiber.StatusInternalServerError, "Failed to create wastebin", false, nil)
			res.SetError(useCaseErr)
			return res.Send(c)
		}

		res := res.NewRes(fiber.StatusOK, "WasteBin created successfully", true, wasteBinEntity)
		return res.Send(c)
	}
}
