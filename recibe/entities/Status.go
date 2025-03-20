package entities

type Status struct {
	Sensores Sensores
}

type Sensores struct {
	Air_quality int
	Temperature  float64
	Humidity      float64
}
