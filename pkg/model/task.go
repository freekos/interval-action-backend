package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	IntervalID  uint   `json:"interval_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
