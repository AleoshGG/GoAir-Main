package domain

type ISensors interface {
	RegisterReadings(reading Readings) (uint, error)
	GetAirQualityAVG(id_place int) []AirQuialityAVG
	GetAirQualityLast24(id_place int) []AirQuialityLast24
	GetTemperatureLast24(id_place int) []TemperatureLast24
	GetHumidityLast24(id_place int) []HumidityLast24
}