package entity

import (
	"time"
)

type Report struct {
	ID            string  `json:"id" gorm:"primaryKey"`
	UserID        *string `json:"user_id" gorm:"index"`
	WasteBinID    *string `json:"wastebin_id" gorm:"index"`
	Weight        *string
	RemainingFill *string
	AirQuality    *string
	Address       *string
	Latitude      *string
	Longitude     *string
	Image         *string   `json:"image"`
	Description   *string   `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
