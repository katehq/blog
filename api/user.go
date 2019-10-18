package api

import (
	"blog/models"
	"blog/serializer"
	"github.com/gin-gonic/gin"
)

// UserList return user
func UserList(c *gin.Context)  {
	user := models.GetFirstUser()
	res := serializer.BuildUserResponse(user)
	c.JSON(200, res)
}

// PostsList list posts
func PostsList(c *gin.Context)  {
	posts := models.GetAll()
	res := serializer.BuildPostsResponse(posts)
	c.JSON(200, res)
}

// PostList post
func PostList(c *gin.Context, post models.Post)  {
	res := serializer.BuildPostResponse(post)
	c.JSON(200, res)
}