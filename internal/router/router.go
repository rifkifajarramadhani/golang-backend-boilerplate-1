package router

import "github.com/gin-gonic/gin"

func Setup() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	registerUserRoutes(api)

	return r
}