package controllers

import (
	"blog/models"
	"blog/api"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

//Userpost create user
func Userpost(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	passwordOnce := c.PostForm("password_once")
	if password != passwordOnce {
		c.JSON(200, gin.H{
			"status": 0,
			"msg":    "密码错误",
		})
		panic("失败")
	}
	fmt.Println(name, password, passwordOnce)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("密码加密失败")
	}
	password = string(hash)
	user := models.User{
		Name:     name,
		Password: password,
	}
	err = user.AddUser()
	if err != nil {
		models.Migrate()
		c.JSON(200, gin.H{
			"status": 0,
			"msg":    "创建失败",
		})
	}
	c.JSON(200, gin.H{
		"status": 1,
		"msg":    "创建用户成功",
		"user":   user,
	})
}

// Userget get all the users
func Userget(c *gin.Context) {
	api.UserList(c)
}
