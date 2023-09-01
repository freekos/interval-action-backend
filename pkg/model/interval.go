package model

type Interval struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	StartAt   string `json:"started_at"`
	StoppedAt string `json:"stopped_at"`
	EndAt     string `json:"ended_at"`
	Tasks     []uint `json:"tasks"`
}
