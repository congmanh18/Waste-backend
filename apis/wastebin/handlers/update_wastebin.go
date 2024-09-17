package handler

import (
	"context"
	req "smart-waste/apis/wastebin/models"
	"smart-waste/domain/wastebin/entity"
	"smart-waste/pkgs/python"
	"smart-waste/pkgs/res"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (w WasteBinHandler) HandlerUpdateWasteBin() fiber.Handler {
	return func(c *fiber.Ctx) error {

		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)

		defer cancel()
		ID := c.Params("id")

		var updateWasteBinReq = new(req.CreateWasteBinReq)
		if err := c.BodyParser(&updateWasteBinReq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		var wasteBinEntity = entity.WasteBin{
			ID:          ID,
			Weight:      updateWasteBinReq.Weight,
			FilledLevel: updateWasteBinReq.FilledLevel,
			AirQuality:  updateWasteBinReq.AirQuality,
			WaterLevel:  updateWasteBinReq.WaterLevel,
			Address:     updateWasteBinReq.Address,
			Latitude:    updateWasteBinReq.Latitude,
			Longitude:   updateWasteBinReq.Longitude,
		}

		output, _ := python.PassDataGoToPy(*updateWasteBinReq.Weight, *updateWasteBinReq.AirQuality, *updateWasteBinReq.WaterLevel, "7200")

		var useCaseErr = w.UpdateWasteBinUsecase.ExecuteUpdateWasteBin(ctx, wasteBinEntity.ID, &wasteBinEntity)
		if useCaseErr != nil {
			res := res.NewRes(fiber.StatusInternalServerError, useCaseErr.Error(), false, nil)
			return res.Send(c)
		}

		res := res.NewRes(fiber.StatusOK, "Waste bin updated successfully", true, wasteBinEntity)
		return res.Send(c)
	}
}

// // Broadcast data tới tất cả client kết nối
// message := fmt.Sprintf("Waste bin ID %s updated", ID)
// BroadcastToClients(message)
