package usecases

import "API/sensors/domain"

type RegisterReadings struct {
	db domain.ISensors
}

func NewRegisterReadings(db domain.ISensors) *RegisterReadings {
	return &RegisterReadings{db: db}
}

func (uc *RegisterReadings) Run(reading domain.Readings) (uint, error) {
	return uc.db.RegisterReadings(reading)
}