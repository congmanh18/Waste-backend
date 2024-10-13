package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
    // Dữ liệu yêu cầu, ví dụ: dự đoán 10 bước tiếp theo
    requestData := map[string]interface{}{
        "steps": 10,
    }

    // Chuyển đổi dữ liệu thành JSON
    jsonData, err := json.Marshal(requestData)
    if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return
    }

    // Gửi yêu cầu POST đến API Flask
    resp, err := http.Post("http://127.0.0.1:5000/predict", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    defer resp.Body.Close()

    // Đọc kết quả trả về
    var result map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&result)

    // In ra kết quả dự đoán
    fmt.Println("Forecast:", result["forecast"])
}
