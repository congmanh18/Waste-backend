package entity

type WasteBin struct {
	ID          string `gorm:"primaryKey"`
	Weight      *string
	FilledLevel *string
	AirQuality  *string
	WaterLevel  *string
	Address     *string
	Latitude    *string
	Longitude   *string
}
