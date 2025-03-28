package controllers

import (
	"API/sensors/application/services"
	usecases "API/sensors/application/useCases"
	"API/sensors/domain"
	"API/sensors/infrastructure"
	"net/http"

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

var count int = 1

func (rm_c *RegisterMetricsController) RegisterMetrics(c *gin.Context) {
	var metrics domain.Sensor
	var newReading domain.Readings

	if err := c.ShouldBindJSON(&metrics); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}

	rm_c.service.Run(metrics)
	count++

	if count == 20 {
		//Creando insert para la base de datos
		newReading.Id_sensor = metrics.Id_sensor[0]
		newReading.Sensor_type = "air_quality"
		newReading.Value = float64(metrics.Air_quality)

		if _, err := rm_c.app.Run(newReading); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": false,
				"error": "Ocurrió un error al insertar en la base de datos en el sensor air_quality: " + err.Error(),
			})
			return
		}

		newReading.Id_sensor = metrics.Id_sensor[1]
		newReading.Sensor_type = "temperature"
		newReading.Value = float64(metrics.Temperature)

		if _, err := rm_c.app.Run(newReading); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": false,
				"error": "Ocurrió un error al insertar en la base de datos en el sensor temperature: " + err.Error(),
			})
			return
		}

		newReading.Id_sensor = metrics.Id_sensor[2]
		newReading.Sensor_type = "humidity"
		newReading.Value = float64(metrics.Humidity)

		if _, err := rm_c.app.Run(newReading); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": false,
				"error": "Ocurrió un error al insertar en la base de datos en el sensor humidity: " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status": true,
			"count": count,
		})
		count = 0
	}	

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"count": count,
	})
}
