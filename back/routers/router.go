package routers

import (
	"back/controllers"
	"back/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://github.com"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(gin.Logger(), cors.New(config))

	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	// Protected routes
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware)
	{
		api.POST("/posts", controllers.CreatePost)
		api.GET("/posts", controllers.GetPosts)
		api.GET("/posts/:id", controllers.GetPost)
		api.PUT("/posts/:id", controllers.UpdatePost)
		api.DELETE("/posts/:id", controllers.DeletePost)
	}

	return router
}
