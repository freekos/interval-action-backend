package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	IntervalID  uint
	Title       string
	Description string
}

type TaskResponse struct {
	ID          uint   `json:"id"`
	IntervalID  uint   `json:"interval_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
