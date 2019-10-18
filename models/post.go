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

// GetPost get single post with id
func (p *Post) GetPost(id int)  {
	db := Open()
	db.Where("id = ?", id).Find(&p)
}

// UpdateViews add the views
func (p *Post) UpdateViews()  {
	db := Open()
	db.Model(p).Update("view", p.View+1)
}