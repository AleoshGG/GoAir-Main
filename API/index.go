package main

import (
	"API/database/conn"
	"API/sensors/infrastructure"
	"API/sensors/infrastructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	conn.Migration()
	

	infrastructure.GoDependences()
	r := gin.Default()
	r.Use(cors.Default())
	
	routes.RegisterRouter(r)
	r.Run()
}