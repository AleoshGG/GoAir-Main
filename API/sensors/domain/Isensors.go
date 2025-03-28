package domain

type ISensors interface {
	RegisterReadings(reading Readings) (uint, error)
	GetMetrics(id_sensor string, sensor_type string) []Readings
}