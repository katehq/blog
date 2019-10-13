package main

import (
	"blog/controllers"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.GET("/posts/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"main": "main",
		})
	})
	router.POST("/users", controllers.Userpost)
	router.GET("/users", controllers.Userget)
	router.POST("/posts/new", controllers.Postpost)
	router.GET("/posts", controllers.Posts)
	router.Run(":8082")
}
