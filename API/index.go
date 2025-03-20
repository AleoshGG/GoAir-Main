package main

import (
	"API/sensors/infrastructure"
	"API/sensors/infrastructure/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	infrastructure.GoDependences()

	r := gin.Default()

	routes.RegisterRouter(r)
	r.Run()
}