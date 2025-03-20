package repositories

import "API/sensors/domain"

type IMetrics interface {
	SendMetrics(metrics domain.Sensor)
}