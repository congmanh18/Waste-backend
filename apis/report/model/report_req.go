package model

type ReportReq struct {
	ID          string  `json:"id"`
	UserID      *string `json:"user_id"`
	WasteBinID  *string `json:"wastebin_id"`
	Image       *string `json:"image"`
	Description *string `json:"description"`
}
