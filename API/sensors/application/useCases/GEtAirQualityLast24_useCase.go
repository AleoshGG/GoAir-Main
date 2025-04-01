package usecases

import "API/sensors/domain"

type GetAirQualityLast24 struct {
	db domain.ISensors
}

func NewGetAirQualityLast24(db domain.ISensors) *GetAirQualityLast24 {
	return &GetAirQualityLast24{db: db}
}

func (uc *GetAirQualityLast24) Run(id_place int) []domain.AirQuialityLast24 {
	return uc.db.GetAirQualityLast24(id_place)
}