package main

import (
	"API/database/conn"
	"API/sensors/infrastructure"
	"API/sensors/infrastructure/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	conn.Connection()
	infrastructure.GoDependences()

	r := gin.Default()

	routes.RegisterRouter(r)
	r.Run()
}