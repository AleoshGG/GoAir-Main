package services

import (
	"API/sensors/application/repositories"
	"API/sensors/domain"
)

type SendMetrics struct {
	mt repositories.IMetrics
}

func NewSendMetricsService(mt repositories.IMetrics) *SendMetrics {
	return &SendMetrics{mt: mt}
}

func (s *SendMetrics) Run(metrics domain.Sensor) {
	s.mt.SendMetrics(metrics)
}