package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	req "smart-waste/apis/wastebin/models"
	"smart-waste/domain/wastebin/entity"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MQTT Client handler for updating WasteBin
func (w WasteBinHandler) MQTTUpdateWasteBin(client mqtt.Client) {
	// Define the MQTT topic for wastebin updates
	topic := "wastebin/update"

	// Set the callback function to be triggered on message arrival
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message on topic: %s\n", topic)

		// Unmarshal the incoming payload (JSON) to CreateWasteBinReq struct
		var updateWasteBinReq req.CreateWasteBinReq
		if err := json.Unmarshal(msg.Payload(), &updateWasteBinReq); err != nil {
			fmt.Printf("Error unmarshaling payload: %v\n", err)
			return
		}

		// Extract ID from the payload (assuming ESP32 sends the bin ID)
		ID := updateWasteBinReq.ID

		// Use case to read the current state of the wastebin from the database
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		wasteBinDB, err := w.ReadWasteBinUsecase.ReadWasteBinByID(ctx, *ID)
		if err != nil {
			fmt.Printf("Unable to load wastebin information for ID %s: %v\n", *ID, err)
			return
		}

		currentTime := time.Now()

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

		// Check if Timestamp exists in wasteBinDB
		if wasteBinDB.Timestamp.IsZero() {
			fmt.Printf("Timestamp is not set in wasteBinDB\n")
			return
		}

		// Calculate Estimated Time to Full
		output, err := EstimatedTimeToFull(wasteBinDB.Timestamp, currentTime, *wasteBinDB.RemainingFill, *updateWasteBinReq.RemainingFill)
		if err != nil {
			fmt.Printf("Error calculating time to full: %v\n", err)
			return
		}

		// Convert RemainingFill to float64 for further calculations
		filledLevel, err := strconv.ParseFloat(*updateWasteBinReq.RemainingFill, 64)
		if err != nil {
			fmt.Printf("Error converting RemainingFill ('%s') to float64: %v\n", *updateWasteBinReq.RemainingFill, err)
			return
		}

		// Predict Time Until Full
		day, hour, minute, second, err := predictTimeUntilFull(filledLevel, output)
		if err != nil {
			fmt.Printf("Error predicting time until full: %v\n", err)
			return
		}

		// Update the waste bin entity with the calculated values
		wasteBinEntity.Day = &day
		wasteBinEntity.Hour = &hour
		wasteBinEntity.Minute = &minute
		wasteBinEntity.Second = &second
		wasteBinEntity.Timestamp = currentTime

		// Execute update in the database
		useCaseErr := w.UpdateWasteBinUsecase.ExecuteUpdateWasteBin(ctx, wasteBinEntity.ID, &wasteBinEntity)
		if useCaseErr != nil {
			fmt.Printf("Error updating waste bin: %v\n", useCaseErr)
			return
		}

		fmt.Printf("Waste bin ID %s updated successfully\n", *ID)
	})
}
