package main

import (
	"encoding/json"
	"log"
	"recibed/entities"
	"recibed/src"

	"github.com/joho/godotenv"
)

func main() {
  // Cargar las variables de entorno
  godotenv.Load()
  rabbit := src.NewRabbitMQ()
  
  // Tratamiento de un mensaje
  msgs := rabbit.GetMessages()
  var forever chan struct{}

  go func() {
    for d := range msgs {
        var status entities.Status
        err := json.Unmarshal(d.Body, &status)
        if err != nil {
            log.Printf("Error al decodificar el mensaje: %s", err)
            continue
        }
        log.Printf(" [x] Prestamo recibido: Calidad del aire: %d%%, Temperatura: %.1f°C, Humedad: %.1f%%", status.Sensores.Calidad_aire, status.Sensores.Temperatura, status.Sensores.Humedad)
    }
}()

  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
  <-forever
}


