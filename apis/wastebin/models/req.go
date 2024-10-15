package req

type CreateWasteBinReq struct {
	ID            *string `json:"id"`
	Weight        *string `json:"weight"`
	RemainingFill *string `json:"remaining_fill"`
	AirQuality    *string `json:"air_quality"`
	WaterLevel    *string `json:"water_level"`
	Address       *string `json:"address"`
	Latitude      *string `json:"latitude"`
	Longitude     *string `json:"longitude"`
}
