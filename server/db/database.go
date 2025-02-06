package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
	gorm.Model
}

type Run struct {
	ID       uint
	UserID   uint
	Name     string
	Distance uint
	Score    uint
	gorm.Model
}

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("running.db"), &gorm.Config{})
	if err != nil {
		panic("failed to open database")
	}
	return db
}
