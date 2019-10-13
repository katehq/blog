package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //sqlite
)

// User user
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

//AddUser add user
func (user *User) AddUser() error {
	db := Open()
	err := db.Create(user).Error
	return err
}

//AllUsers get all users
func AllUsers() []User {
	var users []User
	db := Open()
	db.Find(&users)
	return users
}

//GetFirstUser get it
func GetFirstUser() User {
	var user User
	db := Open()
	db.First(&user)
	return user
}