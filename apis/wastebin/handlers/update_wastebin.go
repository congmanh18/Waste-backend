package handler

import (
	"context"
	"fmt"
	"math"
	req "smart-waste/apis/wastebin/models"
	"smart-waste/domain/wastebin/entity"
	"smart-waste/pkgs/python"
	"smart-waste/pkgs/res"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (w WasteBinHandler) HandlerUpdateWasteBin() fiber.Handler {
	return func(c *fiber.Ctx) error {

		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)

		defer cancel()
		currentTime := time.Now()

		var updateWasteBinReq = new(req.CreateWasteBinReq)
		if err := c.BodyParser(&updateWasteBinReq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		ID := c.Params("id")
		wateBinDB, err := w.ReadWasteBinUsecase.ReadWasteBinByID(ctx, ID)
		if err != nil {
			res := res.NewRes(fiber.StatusNotFound, "Unable to load wastebin information", false, nil)
			res.SetError(err)
			return res.Send(c)
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

		// Check if Timestamp exists in wasteBinEntity
		if wateBinDB.Timestamp.IsZero() {
			res := res.NewRes(fiber.StatusBadRequest, "Timestamp is not set in wasteBinEntity", false, wasteBinEntity)
			return res.Send(c)
		}

		timeSinceStart := currentTime.Sub(wateBinDB.Timestamp).Seconds()

		output, _ := python.PassDataGoToPy(
			*updateWasteBinReq.Weight,
			*updateWasteBinReq.AirQuality,
			*updateWasteBinReq.WaterLevel,
			fmt.Sprintf("%.2f", timeSinceStart)) // Limit the decimal places

		// xem lại
		filledLevel, err := strconv.ParseFloat(*updateWasteBinReq.FilledLevel, 64)
		if err != nil {
			res := res.NewRes(
				fiber.StatusBadRequest,
				fmt.Sprintf("Error converting FilledLevel ('%s') to float64: %v", *updateWasteBinReq.FilledLevel, err),
				false,
				nil)
			return res.Send(c)
		}
		outputFloat, err := strconv.ParseFloat(output, 64)
		if err != nil {
			res := res.NewRes(
				fiber.StatusBadRequest,
				fmt.Sprintf("Error converting output ('%s') to float64: %v", output, err),
				false,
				nil)
			return res.Send(c)
		}
		// để tạm
		day, hour, minute, second, _ := predictTimeUntilFull(filledLevel, outputFloat)

		wasteBinEntity.Day = &day
		wasteBinEntity.Hour = &hour
		wasteBinEntity.Minute = &minute
		wasteBinEntity.Second = &second
		wasteBinEntity.Timestamp = currentTime

		var useCaseErr = w.UpdateWasteBinUsecase.ExecuteUpdateWasteBin(ctx, wasteBinEntity.ID, &wasteBinEntity)
		if useCaseErr != nil {
			res := res.NewRes(fiber.StatusInternalServerError, useCaseErr.Error(), false, nil)
			return res.Send(c)
		}

		res := res.NewRes(fiber.StatusOK, "Waste bin updated successfully", true, wasteBinEntity)
		return res.Send(c)
	}
}

// Temporary function for predicting time until full
func predictTimeUntilFull(filledLevel, predictedRateOfChange float64) (int, int, int, int, error) {

	// Phần trăm còn trông của thùng
	percentRemaining := 100.0 - filledLevel

	if predictedRateOfChange <= 0 {
		return 0, 0, 0, 0, nil
	}

	// Phần trăm còn trống / tỷ lệ thay đổi của thùng
	timeRemainingSeconds := percentRemaining / predictedRateOfChange

	days := int(math.Floor(timeRemainingSeconds / (24 * 3600)))
	hours := int(math.Floor(math.Mod(timeRemainingSeconds, 24*3600) / 3600))
	minutes := int(math.Floor(math.Mod(timeRemainingSeconds, 3600) / 60))
	seconds := int(math.Mod(timeRemainingSeconds, 60))

	return days, hours, minutes, seconds, nil
}

// // Broadcast data tới tất cả client kết nối
// message := fmt.Sprintf("Waste bin ID %s updated", ID)
// BroadcastToClients(message)
