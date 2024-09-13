package entity

import (
	userEntity "smart-waste/domain/user/entity"
	wasteBinEntity "smart-waste/domain/wastebin/entity"
	"time"
)

type Report struct {
	ID          string                  `json:"id" gorm:"primaryKey"`
	User        userEntity.User         `gorm:"foreignKey:UserID"`
	UserID      *string                 `json:"user_id" gorm:"index"`
	WasteBin    wasteBinEntity.WasteBin `gorm:"foreignKey:WasteBinID"`
	WasteBinID  *string                 `json:"wastebin_id" gorm:"index"`
	Image       *string                 `json:"image"`
	Description *string                 `json:"description"`
	CreatedAt   time.Time               `json:"created_at"`
	UpdatedAt   time.Time               `json:"updated_at"`
}
