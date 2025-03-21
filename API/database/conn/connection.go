package conn

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Importa driver de PostgreSQL
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Connection() {
	dbURL := "postgres://postgres:alexis117092@localhost:5432/prueba?sslmode=disable&search_path=metrics"

	m, err := migrate.New(
		"file://./database/migrations",
		dbURL,
	)
	if err != nil {
		log.Fatal("Error al crear el migrador:", err)
	}

	// Aplicar migraciones
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Error al aplicar migraciones:", err)
	}

	fmt.Println("Migraciones aplicadas correctamente")
}
