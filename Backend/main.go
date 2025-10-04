package main

import (
	"Backend/models"
	"Backend/controllers"
	"Backend/config"
	"github.com/gin-gonic/gin"
)

func main() {
	//crear instancia del servidor
	server := gin.Default()

	// Configurar CORS para las solicitudes HTTP desde el frontend
	server.Use(config.SetupCORS())

	//Conectar a la base de datos
	models.ConnectDb()

	//Rutas
	controllers.SetupRoutes(server)
	//iniciar el servidor
	server.Run(":8080")
}
