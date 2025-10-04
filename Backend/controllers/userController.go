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
	var newUsers []models.User
	if err:=c.ShouldBindJSON(&newUsers); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}
	err:= models.InsertManyUsers(newUsers)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}
	c.JSON(http.StatusCreated,gin.H{"message":"usuarios ingresados con exito"})
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
	userid := c.Param("userid")
	user, err := models.GetUserByID(userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	userid := c.Param("userid")
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := models.UpdateUser(userid, updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "usuario actualizado con exito"})
}
func DeleteUser(c *gin.Context) {
	userid := c.Param("userid")
	err := models.DeleteUser(userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "usuario eliminado con exito"})
}
