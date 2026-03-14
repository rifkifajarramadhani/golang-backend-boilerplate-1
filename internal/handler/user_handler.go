package handler

import "github.com/gin-gonic/gin"

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "users get",
	})
}

func GetUserByID(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "users get by id",
	})
}

func CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "users create",
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "users update",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "users delete",
	})
}