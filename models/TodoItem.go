package models

import (
	"time"
)

type TodoItem struct {
  ID uint `gorm:"primaryKey" json:"id"`
	Description string `json:"description"`
	Done bool `json:"done"`
  CreatedAt time.Time `json:"createdAt"`
  UpdatedAt time.Time `json:"updatedAt"`
}