package domain

type SensorType string

const (
	AirQuality  SensorType = "air_quality"
	Temperature SensorType = "temperature"
	Humidity    SensorType = "humidity"
)

type Readings struct {
	Id_sensor   string
	Sensor_type SensorType
	Create_at   string
	Value       float64
}