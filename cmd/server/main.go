package main

import (
	"log"
	"os"
	"task-manager/internal/config"
	"task-manager/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Inicializar base de datos
	config.InitDB()

	// Configurar Gin mode
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Inicializar router
	router := gin.Default()

	// Configurar rutas
	routes.SetupRoutes(router)

	// Obtener puerto del entorno o usar default
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("🚀 Servidor iniciado en puerto %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
