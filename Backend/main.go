package main

import (
	"Backend/models"
	"github.com/gin-gonic/gin"
)

func main() {
	//crear instancia del servidor
	server := gin.Default()

	//Conectar a la base de datos
	models.ConnectDb()

	//Rutas

	//iniciar el servidor
	server.Run(":" + models.PORT)
}
