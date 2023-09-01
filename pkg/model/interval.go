package model

import "gorm.io/gorm"

type Interval struct {
	gorm.Model
	StartAt   string `json:"started_at"`
	StoppedAt string `json:"stopped_at"`
	EndAt     string `json:"ended_at"`
	Tasks     []Task `json:"tasks"`
}
