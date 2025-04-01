package usecases

import "API/sensors/domain"

type GetHumidityLast24 struct {
	db domain.ISensors
}

func NewGetHumidityLast24(db domain.ISensors) *GetHumidityLast24 {
	return &GetHumidityLast24{db: db}
}

func (uc *GetHumidityLast24) Run(id_place int) []domain.HumidityLast24 {
	return uc.db.GetHumidityLast24(id_place)
}