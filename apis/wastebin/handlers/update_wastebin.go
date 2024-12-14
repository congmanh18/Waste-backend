package handler

import (
	"context"
	"fmt"
	"math"
	req "smart-waste/apis/wastebin/models"
	"smart-waste/domain/wastebin/entity"
	"smart-waste/pkgs/res"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Biến toàn cục hoặc được lưu trữ trong một service
var historicalData []float64

const maxHistorySize = 10 // Giới hạn số lượng dữ liệu lịch sử

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
		_, err := w.ReadWasteBinUsecase.ReadWasteBinByID(ctx, ID)
		if err != nil {
			res := res.NewRes(fiber.StatusNotFound, "Unable to load wastebin information", false, nil)
			res.SetError(err)
			return res.Send(c)
		}

		remainingFill, err := strconv.ParseFloat(*updateWasteBinReq.RemainingFill, 64)
		if err != nil {
			res := res.NewRes(
				fiber.StatusBadRequest,
				fmt.Sprintf("Error converting FilledLevel ('%s') to float64: %v", *updateWasteBinReq.RemainingFill, err),
				false,
				nil)
			return res.Send(c)
		}
		filledLevel := (100 - remainingFill)
		// Cập nhật dữ liệu lịch sử
		if len(historicalData) == 0 {
			// Khởi tạo dữ liệu lịch sử nếu chưa có
			historicalData = generateRandomData(filledLevel, maxHistorySize)
		} else {
			// Thêm giá trị mới và loại bỏ giá trị cũ nhất nếu cần
			historicalData = append(historicalData, filledLevel)
			if len(historicalData) > maxHistorySize {
				historicalData = historicalData[1:] // Loại bỏ phần tử đầu tiên
			}
		}

		fmt.Println("hist", historicalData)

		// Sử dụng dữ liệu lịch sử để dự đoán
		day, hour, minute, second, _ := predictWithHistoricalData(filledLevel, historicalData)
		// if err != nil {
		// 	res := res.NewRes(fiber.StatusBadRequest, fmt.Sprintf("Error predicting time until full: %v", err), false, nil)
		// 	return res.Send(c)
		// }

		var wasteBinEntity = entity.WasteBin{
			ID:            ID,
			Weight:        updateWasteBinReq.Weight,
			RemainingFill: updateWasteBinReq.RemainingFill,
			AirQuality:    updateWasteBinReq.AirQuality,
			Address:       updateWasteBinReq.Address,
			Latitude:      updateWasteBinReq.Latitude,
			Longitude:     updateWasteBinReq.Longitude,
			Day:           &day,
			Hour:          &hour,
			Minute:        &minute,
			Second:        &second,
			Timestamp:     currentTime,
		}

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
func predictTimeUntilFullImproved(filledLevel, predictedRateOfChange float64) (int, int, int, int, error) {
	// Kiểm tra dữ liệu đầu vào
	if filledLevel < 0 || filledLevel > 100 {
		return 0, 0, 0, 0, fmt.Errorf("filledLevel (%v) must be between 0 and 100", filledLevel)
	}

	if predictedRateOfChange <= 0 {
		return 0, 0, 0, 0, fmt.Errorf("predictedRateOfChange (%v) must be greater than 0", predictedRateOfChange)
	}

	// Phần trăm còn trống của thùng
	percentRemaining := 100.0 - filledLevel

	// Dự đoán thời gian còn lại
	timeRemainingSeconds := percentRemaining / predictedRateOfChange

	// Chuyển đổi sang ngày, giờ, phút, giây
	days := int(math.Floor(timeRemainingSeconds / (24 * 3600)))
	hours := int(math.Floor(math.Mod(timeRemainingSeconds, 24*3600) / 3600))
	minutes := int(math.Floor(math.Mod(timeRemainingSeconds, 3600) / 60))
	seconds := int(math.Mod(timeRemainingSeconds, 60))

	return days, hours, minutes, seconds, nil
}

// Sử dụng ARIMA hoặc hồi quy phi tuyến tính để dự đoán tốc độ thay đổi
func calculatePredictedRateOfChange(data []float64) (float64, error) {
	// Dữ liệu đầu vào: data chứa lịch sử mức độ lấp đầy theo thời gian
	if len(data) < 2 {
		return 0, fmt.Errorf("not enough data points to calculate rate of change")
	}

	// Tính toán tốc độ thay đổi trung bình
	var totalChange float64
	for i := 1; i < len(data); i++ {
		totalChange += math.Abs(data[i] - data[i-1])
	}

	averageRateOfChange := totalChange / float64(len(data)-1)
	return averageRateOfChange, nil
}

// Hàm dự đoán mới sử dụng cả dữ liệu lịch sử
func predictWithHistoricalData(filledLevel float64, historicalData []float64) (int, int, int, int, error) {
	// Dự đoán tốc độ thay đổi dựa trên dữ liệu lịch sử
	predictedRateOfChange, err := calculatePredictedRateOfChange(historicalData)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("error predicting rate of change: %v", err)
	}

	// Sử dụng hàm cải tiến để dự đoán thời gian còn lại
	return predictTimeUntilFullImproved(filledLevel, predictedRateOfChange)
}

// func EstimatedTimeToFull(previousTimestamp, currentTimestamp time.Time, previousRemainingFill, currentRemainingFill string) (float64, error) {
// 	previousFill, err := strconv.ParseFloat(previousRemainingFill, 64)
// 	if err != nil {
// 		return 0, fmt.Errorf("error converting previousRemainingFill to float64: %v", err)
// 	}

// 	currentFill, err := strconv.ParseFloat(currentRemainingFill, 64)
// 	if err != nil {
// 		return 0, fmt.Errorf("error converting currentRemainingFill to float64: %v", err)
// 	}

// 	timeDiff := currentTimestamp.Sub(previousTimestamp).Seconds()
// 	remainingFillDiff := math.Abs(previousFill - currentFill)

// 	if remainingFillDiff == 0 {
// 		return 0, fmt.Errorf("the remaining fill difference is zero, cannot divide by zero")
// 	}

// 	output := timeDiff * (math.Abs(currentFill) / remainingFillDiff)
// 	return output, nil
// }
