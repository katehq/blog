package main

import (
	"blog/controllers"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.Cors())

	router.POST("/users", controllers.Userpost)
	
	router.GET("/post/:id", controllers.PostByID)
	router.GET("/posts", controllers.Posts)
	router.POST("/login", controllers.Login)

	router.Use(middleware.Auth())
	router.GET("/users", controllers.Userget)
	router.POST("/posts/new", controllers.Postpost)
	router.Run(":8082")
}
