package src

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"recibed/entities"
)

func FetchAPI(metrics entities.Sensores) {
	// Construir la URL de la petición PATCH
	URL := "http://localhost:8080/sensors/"  
	jsonBody, _ := json.Marshal(metrics)

	// Crear la petición PATCH con el body JSON
	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("Error creando la petición PATCH: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Ejecutar la petición PATCH
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error al ejecutar la petición PATCH: %v", err)
	}
	defer resp.Body.Close()

	// Verificar el estado de la respuesta
	if resp.StatusCode != http.StatusOK {
		log.Printf("La petición PATCH devolvió el estado: %d", resp.StatusCode)
	} else {
		log.Println("Libro actualizado correctamente mediante PATCH")
	}
}