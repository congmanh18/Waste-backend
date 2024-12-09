package handler

import (
	"smart-waste/apis/wastebin/data"
	"smart-waste/pkgs/res"

	"github.com/gofiber/fiber/v2"
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
func HandlerExponentialSmoothing() fiber.Handler {
	return func(c *fiber.Ctx) error {
		filledLevelData := data.DataFile()
		alpha := 0.3 // Hệ số alpha cho Exponential Smoothing
		forecast := ForecastExponentialSmoothing(filledLevelData, alpha, 5)

		res := res.NewRes(fiber.StatusOK, "Exponential Smoothing: ", true, forecast)
		return res.Send(c)
	}
}
