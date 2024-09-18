package main

import (
	"fmt"
	"math"
)

// Hàm để dự đoán thời gian thùng rác đầy
func predictTimeUntilFull(filledLevel, predictedRateOfChange float64) string {
	// Bước 1: Tính phần trăm còn lại để thùng rác đầy
	percentRemaining := 100.0 - filledLevel

	// Bước 2: Nếu tốc độ thay đổi nhỏ hoặc bằng 0, không thể dự đoán
	if predictedRateOfChange <= 0 {
		return "Tốc độ thay đổi nhỏ hơn 0"
	}

	// Bước 3: Tính số giây còn lại để thùng rác đầy
	timeRemainingSeconds := percentRemaining / predictedRateOfChange

	// Bước 4: Chuyển đổi thời gian từ giây sang giờ, phút, giây
	hours := int(math.Floor(timeRemainingSeconds / 3600))
	minutes := int(math.Floor(math.Mod(timeRemainingSeconds, 3600) / 60))
	seconds := int(math.Mod(timeRemainingSeconds, 60))

	return fmt.Sprintf("Thùng rác sẽ đầy sau %d giờ, %d phút và %d giây", hours, minutes, seconds)
}

func main() {
	// Ví dụ: FilledLevel hiện tại là 67%, PredictedRateOfChange là 0.0011481846868489198
	filledLevel := 5.0
	predictedRateOfChange := 0.51481846868489198

	// In kết quả ra terminal
	fmt.Println(predictTimeUntilFull(filledLevel, predictedRateOfChange))
}

// package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"

// 	"github.com/cdipaolo/goml/regression"
// )

// func main() {
// 	// Mở tập dữ liệu CSV
// 	file, err := os.Open("data.csv")
// 	if err != nil {
// 		log.Fatalf("Error opening file: %v", err)
// 	}
// 	defer file.Close()

// 	// Đọc tập dữ liệu
// 	reader := csv.NewReader(file)
// 	records, err := reader.ReadAll()
// 	if err != nil {
// 		log.Fatalf("Error reading CSV: %v", err)
// 	}

// 	// Tạo slice chứa các điểm dữ liệu
// 	var dataX [][]float64
// 	var dataY []float64

// 	for _, record := range records[1:] { // bỏ qua header
// 		weight, _ := strconv.ParseFloat(record[0], 64)
// 		airQuality, _ := strconv.ParseFloat(record[1], 64)
// 		waterLevel, _ := strconv.ParseFloat(record[2], 64)
// 		timeSinceStart, _ := strconv.ParseFloat(record[3], 64)
// 		filledLevel, _ := strconv.ParseFloat(record[4], 64)

// 		// Thêm điểm dữ liệu vào slice
// 		dataX = append(dataX, []float64{weight, airQuality, waterLevel, timeSinceStart})
// 		dataY = append(dataY, filledLevel)
// 	}

// 	// Tạo mô hình hồi quy tuyến tính
// 	r := regression.NewLinear(dataX, dataY, 0.1) // 0.1 là learning rate

// 	// Huấn luyện mô hình
// 	go r.Run()

// 	// Chờ mô hình huấn luyện hoàn tất
// 	r.Wait()

// 	// In các hệ số của mô hình
// 	fmt.Println("Hệ số của mô hình:", r.Parameters())
// }
