package utils

import (
	"kevindurb/go-do/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var conn *gorm.DB

func GetConnection() *gorm.DB {
	if (conn == nil) {
		conn, _ = gorm.Open(sqlite.Open("./db.sqlite"), &gorm.Config{})

		conn.AutoMigrate(&models.TodoItem{})
	}

	return conn;
}