package entity

import "time"

type WasteBin struct {
	ID          string `gorm:"primaryKey"`
	Weight      *string
	FilledLevel *string
	AirQuality  *string
	WaterLevel  *string
	Address     *string
	Latitude    *string
	Longitude   *string
	Timestamp   time.Time
	Day         *int
	Hour        *int
	Minute      *int
	Second      *int
}
