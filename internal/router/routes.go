package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rifkifajarramadhani/golang-forum-app/internal/handler"
)

func registerUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", handler.GetUsers)
	users.GET("/:id", handler.GetUserByID)
	users.POST("/", handler.CreateUser)
	users.PUT("/:id", handler.UpdateUser)
	users.DELETE("/:id", handler.DeleteUser)
}