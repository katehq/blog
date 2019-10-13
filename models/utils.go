package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //sqlite
)

// Migrate the db
func Migrate() {
	db := Open()
	defer db.Close()
	db.AutoMigrate(&Post{}, &User{})
}

// Open oepn database
func Open() *gorm.DB {
	db, err := gorm.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic("can not open the database")
	}
	return db
}
