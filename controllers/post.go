package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"blog/models"
)

// Postget get post
func Postget(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": "user",
	})
}

// Postpost post
func Postpost(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	post := models.Post{
		Title:title,
		Content: content,
	}
	post.Add()
	c.JSON(200, gin.H{
		"status":1,
	})
}

// Posts all
func Posts(c *gin.Context)  {
	var posts []models.Post
	posts = models.GetAll()
	c.JSON(200, gin.H{
		"posts":posts,
	})	
}