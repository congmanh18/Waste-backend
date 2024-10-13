Tập dữ liệu bạn cung cấp có các cột như sau:

- `Timestamp`: Thời gian thu thập dữ liệu.
- `Weight(kg)`: Trọng lượng rác trong thùng.
- `FilledLevel(%)`: Mức độ đầy của thùng (%).
- `RemainingFill(%)`: Phần trăm còn lại cho đến khi thùng đầy.
- `TimeDiff`: Khoảng thời gian giữa các lần đo.
- `FillRate(%)_per_second`: Tốc độ làm đầy thùng theo giây (phần trăm mỗi giây).
- `EstimatedTimeToFull(s)`: Thời gian ước tính còn lại (giây) cho đến khi thùng đầy.
- `EstimatedTimeToFull(days)`: Thời gian ước tính còn lại (ngày) cho đến khi thùng đầy.

Dựa vào dữ liệu này, ta có thể xây dựng mô hình dự đoán thời gian cho đến khi thùng đầy bằng cách phân tích mối quan hệ giữa tốc độ làm đầy và phần trăm còn lại.

Để làm mô hình dự đoán trong Golang, bạn có thể thực hiện theo các bước sau:

### Bước 1: Chuẩn bị dữ liệu

1. Đọc tập dữ liệu từ CSV.
2. Chuyển đổi các giá trị thời gian thành các đơn vị đo thời gian (giờ, phút, giây).
3. Xử lý các hàng bị thiếu dữ liệu hoặc không hợp lệ.

### Bước 2: Tính toán tốc độ làm đầy

Sử dụng cột `FillRate(%)_per_second` để tính toán thời gian còn lại cho đến khi thùng đầy, từ đó dự đoán thời gian đến khi thùng đầy dựa trên phần trăm còn lại.

### Bước 3: Viết code Golang

Dưới đây là một đoạn code mẫu Golang:

```go
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type TrashData struct {
	Timestamp             time.Time
	FillLevel             float64
	FillRatePerSecond     float64
	RemainingFillPercent  float64
	EstimatedTimeToFull   float64
}

func parseCSV(filePath string) ([]TrashData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []TrashData
	for i, record := range records {
		if i == 0 {
			// Skip header
			continue
		}

		// Parse the fields
		timestamp, err := time.Parse("2006-01-02 15:04:05", record[1])
		if err != nil {
			return nil, err
		}
		fillLevel, _ := strconv.ParseFloat(record[3], 64)
		fillRate, _ := strconv.ParseFloat(record[7], 64)
		remainingFillPercent, _ := strconv.ParseFloat(record[5], 64)
		estimatedTimeToFull, _ := strconv.ParseFloat(record[8], 64)

		data = append(data, TrashData{
			Timestamp:            timestamp,
			FillLevel:            fillLevel,
			FillRatePerSecond:    fillRate,
			RemainingFillPercent: remainingFillPercent,
			EstimatedTimeToFull:  estimatedTimeToFull,
		})
	}

	return data, nil
}

func predictTimeToFull(trashData TrashData) time.Duration {
	remainingFill := trashData.RemainingFillPercent / 100.0
	fillRate := trashData.FillRatePerSecond

	if fillRate <= 0 {
		return 0 // Invalid or negative fill rate
	}

	// Time to full in seconds
	secondsToFull := remainingFill / fillRate

	// Convert to duration
	durationToFull := time.Duration(secondsToFull) * time.Second
	return durationToFull
}

func main() {
	data, err := parseCSV("Processed_Trash_Fill_Data.csv")
	if err != nil {
		log.Fatal(err)
	}

	for _, trashData := range data {
		timeToFull := predictTimeToFull(trashData)
		fmt.Printf("At %s, time to full is %s\n", trashData.Timestamp, timeToFull)
	}
}
```

### Bước 4: Chạy và kiểm tra
Sau khi viết xong, bạn có thể chạy mô hình dự đoán bằng cách sử dụng `go run` để kiểm tra thời gian đến khi thùng đầy dựa trên tập dữ liệu được cung cấp.

Bạn có thể thay thế phần đọc file với tập dữ liệu thực tế của mình và tinh chỉnh logic tùy theo yêu cầu.

Dưới đây là các công thức để tính từng cột mà bạn yêu cầu:

### 1. **RemainingFill(%)**:
Đây là phần trăm còn lại cho đến khi thùng đầy. Ta có thể tính như sau:

\[
\text{RemainingFill(\%)} = 100\% - \text{FilledLevel(\%)}
\]

### 2. **TimeDiff**:
Khoảng thời gian giữa hai lần đo, ta có thể tính bằng cách lấy sự khác biệt giữa hai lần ghi nhận `Timestamp` liên tiếp.

\[
\text{TimeDiff} = \text{Timestamp}_{i+1} - \text{Timestamp}_i
\]

Nếu bạn có dữ liệu thời gian ở định dạng `datetime`, thì đơn giản là lấy sự khác biệt giữa hai thời điểm liên tiếp.

### 3. **FillRate(%)_per_second**:
Tốc độ làm đầy thùng tính theo giây (phần trăm mỗi giây). Tính dựa trên sự thay đổi mức độ đầy của thùng qua một khoảng thời gian (`TimeDiff`).

\[
\text{FillRate(\%) per second} = \frac{\text{FilledLevel}_{i+1} - \text{FilledLevel}_i}{\text{TimeDiff (giây)}}
\]

Ví dụ, nếu mức độ đầy tăng từ 10% lên 20% trong 30 phút (1800 giây), thì tốc độ làm đầy sẽ là:

\[
\text{FillRate} = \frac{20\% - 10\%}{1800 \text{ giây}} = 0.00556\%/\text{giây}
\]

### 4. **EstimatedTimeToFull(s)**:
Thời gian ước tính cho đến khi thùng đầy (tính bằng giây) có thể được tính bằng cách chia `RemainingFill(%)` cho tốc độ làm đầy (`FillRate(%)_per_second`).

\[
\text{EstimatedTimeToFull(s)} = \frac{\text{RemainingFill(\%)}}{\text{FillRate(\%) per second}}
\]

Nếu tốc độ làm đầy là 0.00556% mỗi giây và phần trăm còn lại là 80%, thì thời gian còn lại cho đến khi thùng đầy sẽ là:

\[
\text{EstimatedTimeToFull(s)} = \frac{80\%}{0.00556\%/\text{giây}} = 14423 \text{ giây}
\]

### 5. **EstimatedTimeToFull(days)**:
Chuyển đổi từ giây sang ngày:

\[
\text{EstimatedTimeToFull(days)} = \frac{\text{EstimatedTimeToFull(s)}}{86400}
\]

(86400 giây trong một ngày).

Ví dụ: nếu `EstimatedTimeToFull(s)` là 14423 giây, thì thời gian còn lại tính bằng ngày là:

\[
\text{EstimatedTimeToFull(days)} = \frac{14423}{86400} = 0.167 \text{ ngày} = khoảng 4 giờ.
\]

### Tổng kết:

- **RemainingFill(%):** `100% - FilledLevel(%)`.
- **TimeDiff:** Chênh lệch giữa hai mốc thời gian.
- **FillRate(%) per second:** Tốc độ thay đổi phần trăm mỗi giây, tính bằng \(\frac{\Delta \text{FilledLevel(\%)}}{\text{TimeDiff}}\).
- **EstimatedTimeToFull(s):** Thời gian còn lại cho đến khi đầy thùng, tính bằng \(\frac{\text{RemainingFill(\%)}}{\text{FillRate(%) per second}}\).
- **EstimatedTimeToFull(days):** Thời gian còn lại cho đến khi đầy thùng, tính bằng \(\frac{\text{EstimatedTimeToFull(s)}}{86400}\).

Những công thức này sẽ giúp bạn tính toán các giá trị cần thiết từ dữ liệu thu thập.