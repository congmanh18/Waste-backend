package data

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func DataFile() []float64 {
	// Đường dẫn đến tệp CSV local (ví dụ: tệp đã được sao chép vào Docker image)
	filePath := "/dist/data/Processed_Trash_Fill_Data.csv"

	// Mở tệp CSV từ local
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Đọc dữ liệu CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV data: %v", err)
	}

	// Xử lý dữ liệu CSV, lấy các giá trị từ cột "FilledLevel"
	var filledLevelData []float64
	for _, record := range records[1:] { // Bỏ qua dòng tiêu đề
		filledLevel := record[2] // Giả sử cột thứ 2 là "FilledLevel"
		value, err := strconv.ParseFloat(filledLevel, 64)
		if err != nil {
			log.Printf("Error parsing FilledLevel value: %v", err)
			continue
		}
		filledLevelData = append(filledLevelData, value)
	}

	return filledLevelData
}
