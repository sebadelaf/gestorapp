package controllers

import "github.com/gin-gonic/gin"

func SetupRoutes(server *gin.Engine) {
	//Rutas
	users := server.Group("/user")
	{   // la estructura es /user + lo que venga despues, metodo del controller
		users.GET("/", GetUsers)
		users.GET("/:userid", GetUserByID)
		users.POST("/new", CreateUser)
		users.POST("/newmany", CreateManyUsers)
		users.PUT("/:userid", UpdateUser)
		users.DELETE("/:userid", DeleteUser)
	}
}
