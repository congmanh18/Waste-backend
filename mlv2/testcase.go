package mlv2

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/cors"
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
func ForecastHandler(w http.ResponseWriter, r *http.Request) {
	// Dữ liệu thùng rác (có thể thay thế bằng dữ liệu từ cơ sở dữ liệu)
	filledLevelData := []float64{3.05, 5.25, 9.19, 14.18, 14.18, 14.18, 14.18, 14.18, 23.52, 31.20, 39.60, 49.35}
	alpha := 0.3 // Hệ số alpha cho Exponential Smoothing
	forecast := ForecastExponentialSmoothing(filledLevelData, alpha, 5)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(forecast)
}

func TestCase() {
	// Cấu hình CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Cho phép tất cả các origin, bạn có thể thay "*" bằng "http://localhost:3000" nếu muốn giới hạn
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Đăng ký handler cho endpoint
	http.HandleFunc("/forecast", ForecastHandler)

	// Khởi chạy server ở cổng 8080 và sử dụng middleware CORS
	fmt.Println("Server đang chạy ở cổng 8080...")
	http.ListenAndServe(":8080", corsHandler.Handler(http.DefaultServeMux))
}
