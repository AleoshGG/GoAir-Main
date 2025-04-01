package usecases

import "API/sensors/domain"

type GetAirQualityAVG struct {
	db domain.ISensors
}

func NewGetAirQualityAVG(db domain.ISensors) *GetAirQualityAVG {
	return &GetAirQualityAVG{db: db}
}

func (uc *GetAirQualityAVG) Run(id_place int) []domain.AirQuialityAVG {
	return uc.db.GetAirQualityAVG(id_place)
}