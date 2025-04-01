package usecases

import "API/sensors/domain"

type GetTemperatureLast24 struct {
	db domain.ISensors
}

func NewGetTemperatureLast24(db domain.ISensors) *GetTemperatureLast24 {
	return &GetTemperatureLast24{db: db}
}

func (uc *GetTemperatureLast24) Run(id_place int) []domain.TemperatureLast24 {
	return uc.db.GetTemperatureLast24(id_place)
}