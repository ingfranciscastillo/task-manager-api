package config

import (
	"fmt"
	"log"
	"os"
	"task-manager/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB inicializa la conexión a la base de datos
func InitDB() {
	var err error

	// Construir DSN (Data Source Name) para PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	// Conectar a la base de datos
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	// Ejecutar migraciones automáticas
	err = DB.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		log.Fatal("Error al ejecutar migraciones:", err)
	}

	log.Println("✅ Base de datos conectada exitosamente")
}

// GetDB retorna la instancia de la base de datos
func GetDB() *gorm.DB {
	return DB
}
