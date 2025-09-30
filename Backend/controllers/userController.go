package controllers

import (
	"github.com/gin-gonic/gin"
	"Backend/models"
	"net/http"
)

func CreateUser(c *gin.Context) {
	//creamos la variable donde se guardara el valor dado
	var newUser models.User
	if err:=c.ShouldBindJSON(&newUser); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}
	err:= models.InsertUser(newUser)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}
	c.JSON(http.StatusCreated,gin.H{"message":"usuario ingresado con exito"})
}

func CreateManyUsers(c *gin.Context) {

}

func GetUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
func GetUserByID(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}
func DeleteUser(c *gin.Context) {

}
