package model

import "gorm.io/gorm"

type Interval struct {
	gorm.Model
	StartAt   string
	StoppedAt string
	EndAt     string
}

type IntervalResponse struct {
	ID        uint           `json:"id"`
	StartAt   string         `json:"start_at"`
	StoppedAt string         `json:"stopped_at"`
	EndAt     string         `json:"end_at"`
	Tasks     []TaskResponse `json:"tasks"`
}
