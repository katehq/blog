package controllers

import (
	"blog/api"
	"blog/models"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

// Login login
func Login(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	db := models.Open()
	var user models.User
	db.Where("name=?", name).Find(&user)
	errnew := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errnew != nil {
		c.JSON(200, gin.H{
			"msg": "Password is wrong",
		})
		panic("password is wrong")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
		"exp":  time.Now().Add(time.Hour * 120).Unix(),
		"iss":  "https://katehq.com",
	})

	secret := []byte("luowanshun")
	tokenString, errsign := token.SignedString(secret)
	if errsign != nil {
		panic("wrong in errsign")
	}
	c.JSON(200, gin.H{
		"token": tokenString,
	})
}
