package routes

import (
	"github.com/gin-gonic/gin"
	"controllers"
	"net/http"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/token", controllers.Token)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}