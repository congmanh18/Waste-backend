package entity

import "time"

type Report struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	UserID      string    `json:"user_id"`
	WasteBinID  *string   `json:"wastebin_id"`
	Image       *string   `json:"image"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
