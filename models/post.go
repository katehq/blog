package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite
)

// Post the article
type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	View    uint   `json:"view"`
}

// Add a post
func (p *Post) Add() {
	db := Open()
	db.Create(p)
}


// GetAll return all posts
func GetAll() []Post {
	db := Open()
	var posts []Post
	db.Find(&posts)
	return posts
}
