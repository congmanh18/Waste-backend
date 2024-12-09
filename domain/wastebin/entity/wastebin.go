package entity

import "time"

type WasteBin struct {
	ID            string    `gorm:"primaryKey" json:"id"`
	Weight        *string   `json:"weight"`
	RemainingFill *string   `json:"remaining_fill"`
	AirQuality    *string   `json:"air_quality"`
	Address       *string   `json:"address"`
	Latitude      *string   `json:"latitude"`
	Longitude     *string   `json:"longitude"`
	Timestamp     time.Time `json:"timestamp"`
	Day           *int      `json:"day"`
	Hour          *int      `json:"hour"`
	Minute        *int      `json:"minute"`
	Second        *int      `json:"second"`
}
