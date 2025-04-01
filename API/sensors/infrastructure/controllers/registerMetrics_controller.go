package controllers

import (
	"API/sensors/application/services"
	usecases "API/sensors/application/useCases"
	"API/sensors/domain"
	"API/sensors/infrastructure"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

type RegisterMetricsController struct {
	service *services.SendMetrics
	app 	*usecases.RegisterReadings
}

func NewRegisterMetricsController() *RegisterMetricsController {
	rabbit := infrastructure.GetRabbitMQ()
	service := services.NewSendMetricsService(rabbit)
	postgres := infrastructure.GetPostgreSQL()
	app := usecases.NewRegisterReadings(postgres)
	return &RegisterMetricsController{service: service, app: app}
}

var count int32 = 1

func (rm_c *RegisterMetricsController) RegisterMetrics(c *gin.Context) {
	var metrics domain.Sensor

	if err := c.ShouldBindJSON(&metrics); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}

	rm_c.service.Run(metrics)
	current := atomic.AddInt32(&count, 1)

	if current == 20 {
        // Copiar datos necesarios para la goroutine
        data := struct {
            IDs      []string
            Air      int
            Temp     float64
            Humidity float64
        }{
            IDs:      metrics.Id_sensor,
            Air:      metrics.Air_quality,
            Temp:     metrics.Temperature,
            Humidity: metrics.Humidity,
        }

        go func(d struct { IDs []string; Air int; Temp, Humidity float64 }) {
            // Insertar air_quality
            if _, err := rm_c.app.Run(domain.Readings{
                Id_sensor:   d.IDs[0],
                Sensor_type: "air_quality",
                Value:       float64(d.Air),
            }); err != nil {
                log.Printf("Error insertando air_quality: %v", err) // <- Usar logging
            }

            // Insertar temperature
            if _, err := rm_c.app.Run(domain.Readings{
                Id_sensor:   d.IDs[1],
                Sensor_type: "temperature",
                Value:       float64(d.Temp),
            }); err != nil {
                log.Printf("Error insertando temperature: %v", err)
            }

            // Insertar humidity
            if _, err := rm_c.app.Run(domain.Readings{
                Id_sensor:   d.IDs[2],
                Sensor_type: "humidity",
                Value:       float64(d.Humidity),
            }); err != nil {
                log.Printf("Error insertando humidity: %v", err)
            }

            atomic.StoreInt32(&count, 0) // Reset después de inserts
        }(data)
    }

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"count": count,
	})
}
