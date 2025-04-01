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
	
	
	res, err := postgre.conn.DB.Query(query, reading.Id_sensor, reading.Sensor_type, reading.Value)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta 1: %v", err)
		return 0, err
	}
	defer res.Close()
	return 1, nil
} 

func (postgre *PostgreSQL) GetAirQualityAVG(id_place int) []domain.AirQuialityAVG {
	query := "SELECT * FROM get_air_quality_avg($1)";
    var metrics []domain.AirQuialityAVG

	rows, err := postgre.conn.DB.Query(query, id_place)
	if err != nil {
        fmt.Println("No se pudieron obtener los datos.", err)
        return []domain.AirQuialityAVG{}
    }

	defer rows.Close()

	for rows.Next() {
		var a domain.AirQuialityAVG
		
		// Escanear los valores de la fila
		err := rows.Scan(&a.Fecha, &a.Promedio_calidad_aire)
		if err != nil {
			// Manejar error al escanear la fila
			fmt.Println("Error al escanear la fila:", err)
			return []domain.AirQuialityAVG{}
		}
		metrics = append(metrics, a)
	}

	// Verifica errores después de iterar
    if err = rows.Err(); err != nil {
        fmt.Println("Error al recorrer las filas:", err)
        return nil
    }

	return metrics
}

func (postgre *PostgreSQL) GetAirQualityLast24(id_place int) []domain.AirQuialityLast24 {
	query := "SELECT * FROM get_air_quality_last24($1)";
    var metrics []domain.AirQuialityLast24

	rows, err := postgre.conn.DB.Query(query, id_place)
	if err != nil {
        fmt.Println("No se pudieron obtener los datos.", err)
        return []domain.AirQuialityLast24{}
    }

	defer rows.Close()

	for rows.Next() {
		var a domain.AirQuialityLast24
		
		// Escanear los valores de la fila
		err := rows.Scan(&a.Hora, &a.Calidad_promedio)
		if err != nil {
			// Manejar error al escanear la fila
			fmt.Println("Error al escanear la fila:", err)
			return []domain.AirQuialityLast24{}
		}
		metrics = append(metrics, a)
	}

	// Verifica errores después de iterar
    if err = rows.Err(); err != nil {
        fmt.Println("Error al recorrer las filas:", err)
        return nil
    }

	return metrics
}

func (postgre *PostgreSQL) GetTemperatureLast24(id_place int) []domain.TemperatureLast24 {
	query := "SELECT * FROM get_temperature_last24($1)";
    var metrics []domain.TemperatureLast24

	rows, err := postgre.conn.DB.Query(query, id_place)
	if err != nil {
        fmt.Println("No se pudieron obtener los datos.", err)
        return []domain.TemperatureLast24{}
    }

	defer rows.Close()

	for rows.Next() {
		var a domain.TemperatureLast24
		
		// Escanear los valores de la fila
		err := rows.Scan(&a.Hora, &a.Temperatura_promedio)
		if err != nil {
			// Manejar error al escanear la fila
			fmt.Println("Error al escanear la fila:", err)
			return []domain.TemperatureLast24{}
		}
		metrics = append(metrics, a)
	}

	// Verifica errores después de iterar
    if err = rows.Err(); err != nil {
        fmt.Println("Error al recorrer las filas:", err)
        return nil
    }

	return metrics
}

func (postgre *PostgreSQL) GetHumidityLast24(id_place int) []domain.HumidityLast24 {
	query := "SELECT * FROM get_humidity_last24($1)";
    var metrics []domain.HumidityLast24

	rows, err := postgre.conn.DB.Query(query, id_place)
	if err != nil {
        fmt.Println("No se pudieron obtener los datos.", err)
        return []domain.HumidityLast24{}
    }

	defer rows.Close()

	for rows.Next() {
		var a domain.HumidityLast24
		
		// Escanear los valores de la fila
		err := rows.Scan(&a.Hora, &a.Humedad_promedio)
		if err != nil {
			// Manejar error al escanear la fila
			fmt.Println("Error al escanear la fila:", err)
			return []domain.HumidityLast24{}
		}
		metrics = append(metrics, a)
	}

	// Verifica errores después de iterar
    if err = rows.Err(); err != nil {
        fmt.Println("Error al recorrer las filas:", err)
        return nil
    }

	return metrics
}
