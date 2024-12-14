package handler

import (
	"context"
	"smart-waste/pkgs/res"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/rand"
)

// Hàm Exponential Smoothing (Simple Exponential Smoothing)
func ExponentialSmoothing(data []float64, alpha float64) []float64 {
	smoothed := make([]float64, len(data))
	smoothed[0] = data[0] // Bắt đầu với giá trị đầu tiên của dữ liệu

	// Áp dụng công thức Exponential Smoothing cho phần còn lại của dữ liệu
	for i := 1; i < len(data); i++ {
		smoothed[i] = alpha*data[i] + (1-alpha)*smoothed[i-1]
	}

	return smoothed
}

// Hàm dự báo với Exponential Smoothing
func ForecastExponentialSmoothing(data []float64, alpha float64, forecastPeriods int) []float64 {
	smoothed := ExponentialSmoothing(data, alpha)
	lastValue := smoothed[len(smoothed)-1] // Giá trị cuối cùng từ dữ liệu đã làm mượt

	// Dự báo các giá trị trong tương lai
	forecast := make([]float64, forecastPeriods)
	for i := 0; i < forecastPeriods; i++ {
		lastValue = alpha*data[len(data)-1] + (1-alpha)*lastValue
		forecast[i] = lastValue
	}

	return forecast
}

// API endpoint trả về dự báo mức độ đầy thùng rác
func (w WasteBinHandler) HandlerExponentialSmoothing() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
		defer cancel()

		wateBinEntity, err := w.ReadWasteBinUsecase.ReadWasteBinByID(ctx, c.Params("id"))
		if err != nil {
			res := res.NewRes(fiber.StatusNotFound, "Unable to load wastebin information", false, nil)
			res.SetError(err)
			return res.Send(c)
		}

		// Chuyển đổi RemainingFill từ string sang float64
		remainingLevelFromID, err := strconv.ParseFloat(*wateBinEntity.RemainingFill, 64)
		if err != nil {
			res := res.NewRes(fiber.StatusBadRequest, "Invalid fill level data", false, nil)
			res.SetError(err)
			return res.Send(c)
		}

		filledLevelFromID := 100 - remainingLevelFromID
		filledLevelData := generateRandomData(filledLevelFromID, 4)
		filledLevelData = append(filledLevelData, (filledLevelFromID + 4))

		alpha := 0.5 // Hệ số alpha cho Exponential Smoothing
		forecast := ForecastExponentialSmoothing(filledLevelData, alpha, 5)

		res := res.NewRes(fiber.StatusOK, "Exponential Smoothing: ", true, forecast)
		return res.Send(c)
	}
}

func generateRandomData(max float64, count int) []float64 {
	rand.Seed(uint64(time.Now().UnixNano()))

	if max <= 0 {
		return make([]float64, count)
	}

	min := max - 35
	if min < 0 {
		min = 0
	}

	// Khoảng cách giữa các số (cố định)
	step := (max - min) / float64(count)

	data := make([]float64, count)
	data[0] = min + rand.Float64()*(max-min) // Giá trị đầu tiên ngẫu nhiên trong khoảng [min, max]

	// Tạo các giá trị tiếp theo cách nhau một khoảng cố định
	for i := 1; i < count; i++ {
		data[i] = data[i-1] + step + rand.Float64()*step // Thêm một khoảng ngẫu nhiên nhỏ hơn step
		// Đảm bảo giá trị không vượt quá max
		if data[i] > max {
			data[i] = max
		}
	}

	return data
}
