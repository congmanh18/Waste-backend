package req

type CreateWasteBinReq struct {
	Weight      *string `json:"weight"`
	FilledLevel *string `json:"filled_level"`
	AirQuality  *string `json:"air_quality"`
	WaterLevel  *string `json:"water_level"`
	Address     *string `json:"address"`
	Latitude    *string `json:"latitude"`
	Longitude   *string `json:"longitude"`
}
