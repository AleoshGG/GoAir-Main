package controllers

import (
	"API/sensors/application/services"
	"API/sensors/domain"
	"API/sensors/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterMetricsController struct {
	service *services.SendMetrics
}

func NewRegisterMetricsController() *RegisterMetricsController {
	rabbit := infrastructure.GetRabbitMQ()
	service := services.NewSendMetricsService(rabbit)
	return &RegisterMetricsController{service: service}
}

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

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
	})
}
