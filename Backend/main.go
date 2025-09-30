package main

import (
	"Backend/models"
	"github.com/gin-gonic/gin"
	"Backend/controllers"
)

func main() {
	//crear instancia del servidor
	server := gin.Default()


	//Conectar a la base de datos
	models.ConnectDb()

	//Rutas
	users:=server.Group("/user")
	{	users.GET("/", controllers.GetUsers)
		users.POST("/new", controllers.CreateUser)
	}
	//iniciar el servidor
	server.Run()
}
