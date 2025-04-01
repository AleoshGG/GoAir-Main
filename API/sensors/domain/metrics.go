package domain

type AirQuialityAVG struct {
	Fecha                 string
	Promedio_calidad_aire float64
}

type TemperatureLast24 struct {
	Hora                 string
	Temperatura_promedio float64
}

type HumidityLast24 struct {
	Hora             string
	Humedad_promedio float64
}

type AirQuialityLast24 struct {
	Hora             string
	Calidad_promedio float64
}