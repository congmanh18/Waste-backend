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
	defer c.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	currentTime := time.Now()

	var updateWasteBinReq req.CreateWasteBinReq

	// Đọc dữ liệu JSON từ ESP32 qua WebSocket
	if err := c.ReadJSON(&updateWasteBinReq); err != nil {
		c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error reading JSON: %v", err)))
		return
	}

	ID := updateWasteBinReq.ID // Lấy ID từ dữ liệu JSON
	wateBinDB, err := w.ReadWasteBinUsecase.ReadWasteBinByID(ctx, *ID)
	if err != nil {
		message := fmt.Sprintf("Unable to load wastebin information: %v", err)
		c.WriteMessage(websocket.TextMessage, []byte(message))
		return
	}

	var wasteBinEntity = entity.WasteBin{
		ID:            *ID,
		Weight:        updateWasteBinReq.Weight,
		RemainingFill: updateWasteBinReq.RemainingFill,
		AirQuality:    updateWasteBinReq.AirQuality,
		WaterLevel:    updateWasteBinReq.WaterLevel,
		Address:       updateWasteBinReq.Address,
		Latitude:      updateWasteBinReq.Latitude,
		Longitude:     updateWasteBinReq.Longitude,
	}

	// Check if Timestamp exists in wasteBinEntity
	if wateBinDB.Timestamp.IsZero() {
		c.WriteMessage(websocket.TextMessage, []byte("Timestamp is not set in wasteBinEntity"))
		return
	}

	output, err := EstimatedTimeToFull(wateBinDB.Timestamp, currentTime, *wateBinDB.RemainingFill, *updateWasteBinReq.RemainingFill)
	if err != nil {
		c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error calculating time to full: %v", err)))
		return
	}

	// Parse RemainingFill to float
	filledLevel, err := strconv.ParseFloat(*updateWasteBinReq.RemainingFill, 64)
	if err != nil {
		c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error converting FilledLevel: %v", err)))
		return
	}

	day, hour, minute, second, _ := predictTimeUntilFull(filledLevel, output)

	wasteBinEntity.Day = &day
	wasteBinEntity.Hour = &hour
	wasteBinEntity.Minute = &minute
	wasteBinEntity.Second = &second
	wasteBinEntity.Timestamp = currentTime

	// Update WasteBin trong database
	useCaseErr := w.UpdateWasteBinUsecase.ExecuteUpdateWasteBin(ctx, wasteBinEntity.ID, &wasteBinEntity)
	if useCaseErr != nil {
		c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error updating WasteBin: %v", useCaseErr)))
		return
	}

	// Phản hồi lại cho ESP32 qua WebSocket
	c.WriteMessage(websocket.TextMessage, []byte("Waste bin updated successfully"))
}
