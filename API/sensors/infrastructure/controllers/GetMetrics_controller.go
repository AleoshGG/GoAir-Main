package controllers

import (
	usecases "API/sensors/application/useCases"
	"API/sensors/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetMetrics struct {
	a3 *usecases.GetAirQualityAVG
	a24 *usecases.GetAirQualityLast24
	t24 *usecases.GetTemperatureLast24
	h24 *usecases.GetHumidityLast24
}

func NewGetMetrics() *GetMetrics {
	postgres := infrastructure.GetPostgreSQL()
	a3 := usecases.NewGetAirQualityAVG(postgres)
	a24 := usecases.NewGetAirQualityLast24(postgres)
	t24 := usecases.NewGetTemperatureLast24(postgres)
	h24 := usecases.NewGetHumidityLast24(postgres)
	return &GetMetrics{a3: a3, a24: a24, t24: t24, h24: h24}
}

func (gm_c *GetMetrics) GetMetrics(c *gin.Context) {
	id := c.Param("id")
	id_palce, _ := strconv.ParseInt(id, 10, 64)

	airQuialityAVG := gm_c.a3.Run(int(id_palce))
	airQuialitiLast24 := gm_c.a24.Run(int(id_palce))
	tem24 := gm_c.t24.Run(int(id_palce))
	hum24 := gm_c.h24.Run(int(id_palce))	

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/admin/",
		},
		"AirQualityAVG": airQuialityAVG,
		"AirQuality24": airQuialitiLast24,
		"Temperature24": tem24,
		"Humidity24": hum24,
	})
}