package main

import (
	"API/sensors/infrastructure"
	"API/sensors/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	infrastructure.GoDependences()

	r := gin.Default()

	routes.RegisterRouter(r)
	r.Run()
}