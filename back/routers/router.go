package routers

import (
	"back/controllers"

	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Go API!",
		})
	})

	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	return router
}
