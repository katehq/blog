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