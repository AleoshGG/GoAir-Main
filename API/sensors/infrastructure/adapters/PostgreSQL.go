package adapters

import (
	"API/database/conn"
	"API/sensors/domain"
	"fmt"
)

type PostgreSQL struct {
	conn *conn.ConnPostgreSQL
}

func NewPostgreSQL() *PostgreSQL {
	conn := conn.GetDBPool()

	if conn.Err != "" {
		fmt.Println("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &PostgreSQL{conn: conn}
}

func (postgre *PostgreSQL) RegisterReadings(reading domain.Readings) (uint, error) {
	query := "INSERT INTO sensors_readings (id_sensor, sensor_type, value) VALUES ($1,$2,$3)"

	_, err := postgre.conn.DB.Query(query, reading.Id_sensor, reading.Sensor_type, reading.Value)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta 1: %v", err)
		return 0, err
	}

	return 1, nil
} 

func (postgre *PostgreSQL) GetMetrics(id_sensor string, sensor_type string) []domain.Readings {
	return []domain.Readings{}
}