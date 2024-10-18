package handler

import (
	"context"
	"fmt"
	req "smart-waste/apis/wastebin/models"
	"smart-waste/domain/wastebin/entity"
	"strconv"
	"time"

	"github.com/gofiber/websocket/v2"
)

func (w WasteBinHandler) WebSocketUpdateWasteBin(c *websocket.Conn) {
	defer c.Close() // Đóng khi gặp lỗi hoặc khi kết thúc vòng lặp

	// Tạo ticker để gửi ping/heartbeat mỗi 30 giây
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	// Khởi tạo context bên ngoài vòng lặp để không bị đóng sớm
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		for range ticker.C { // Sử dụng for range để lặp qua các tín hiệu từ ticker
			// Gửi ping để giữ kết nối WebSocket sống
			if err := c.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				fmt.Println("Error sending ping:", err)
				return
			}
		}
	}()

	for {
		currentTime := time.Now()

		var updateWasteBinReq req.CreateWasteBinReq

		// Đọc dữ liệu JSON từ ESP32 qua WebSocket
		if err := c.ReadJSON(&updateWasteBinReq); err != nil {
			// Nếu gặp lỗi, gửi thông báo và tiếp tục đọc lại sau 2 giây
			c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error reading JSON: %v", err)))
			time.Sleep(2 * time.Second) // Đợi một thời gian trước khi thử lại
			continue                    // Tiếp tục vòng lặp để đọc lại tin nhắn
		}

		ID := updateWasteBinReq.ID // Lấy ID từ dữ liệu JSON
		wateBinDB, err := w.ReadWasteBinUsecase.ReadWasteBinByID(ctx, *ID)
		if err != nil {
			message := fmt.Sprintf("Unable to load wastebin information: %v", err)
			c.WriteMessage(websocket.TextMessage, []byte(message))
			continue // Không ngắt kết nối, tiếp tục vòng lặp để đọc dữ liệu tiếp theo
		}

		var wasteBinEntity = entity.WasteBin{
			ID:            *ID,
			Weight:        updateWasteBinReq.Weight,
			RemainingFill: updateWasteBinReq.RemainingFill,
			AirQuality:    updateWasteBinReq.AirQuality,
			Address:       updateWasteBinReq.Address,
			Latitude:      updateWasteBinReq.Latitude,
			Longitude:     updateWasteBinReq.Longitude,
		}

		// Kiểm tra Timestamp
		if wateBinDB.Timestamp.IsZero() {
			c.WriteMessage(websocket.TextMessage, []byte("Timestamp is not set in wasteBinEntity"))
			continue // Chuyển sang lắng nghe tin nhắn tiếp theo
		}

		output, _ := EstimatedTimeToFull(wateBinDB.Timestamp, currentTime, *wateBinDB.RemainingFill, *updateWasteBinReq.RemainingFill)

		// Parse RemainingFill to float
		filledLevel, err := strconv.ParseFloat(*updateWasteBinReq.RemainingFill, 64)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error converting FilledLevel: %v", err)))
			continue // Gửi lỗi và lắng nghe tin nhắn tiếp theo
		}

		day, hour, minute, second, _ := predictTimeUntilFull(filledLevel, output)

		wasteBinEntity.Day = &day
		wasteBinEntity.Hour = &hour
		wasteBinEntity.Minute = &minute
		wasteBinEntity.Second = &second
		wasteBinEntity.Timestamp = currentTime

		// Cập nhật WasteBin trong database
		useCaseErr := w.UpdateWasteBinUsecase.ExecuteUpdateWasteBin(ctx, wasteBinEntity.ID, &wasteBinEntity)
		if useCaseErr != nil {
			c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error updating WasteBin: %v", useCaseErr)))
			continue // Gửi lỗi và lắng nghe tin nhắn tiếp theo
		}

		// Phản hồi lại cho ESP32 qua WebSocket
		c.WriteMessage(websocket.TextMessage, []byte("Waste bin updated successfully"))
	}
}
