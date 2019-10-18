package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"blog/models"
	"strconv"
	"blog/api"
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
	api.PostsList(c)	
}

// PostByID find by id
func PostByID(c *gin.Context){
	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		panic("not find")
	}
	var p models.Post
	p.GetPost(idint)
	p.UpdateViews()
	api.PostList(c,p)
}