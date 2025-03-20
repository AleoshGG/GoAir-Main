package routes

import (
	"API/sensors/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	sensorsRoutes := r.Group("/sensors")
	{
		sensorsRoutes.POST("/", controllers.NewRegisterMetricsController().RegisterMetrics)
	}
}