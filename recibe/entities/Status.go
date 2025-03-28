package entities

type Status struct {
	Sensores Sensores
}

type Sensores struct {
	Id_sensor   []string
	Air_quality int
	Temperature float64
	Humidity    float64
	Ventilador  string
}
