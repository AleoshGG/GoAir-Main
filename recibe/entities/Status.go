package entities

type Status struct {
	Sensores Sensores
}

type Sensores struct {
	Calidad_aire int
	Temperatura  float64
	Humedad      float64
}
