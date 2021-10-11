package models

import "gorm.io/gorm"

type TodoItem struct {
	gorm.Model
	Description string `json:"description"`
	Done bool `json:"done"`
}