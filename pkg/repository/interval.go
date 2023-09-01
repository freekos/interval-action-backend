package repository

import (
	"interval-action/pkg/model"
)

func (repo *Repository) CreateInterval(startAt string, endAt string) model.Interval {
	interval := model.Interval{StartAt: startAt, EndAt: endAt}

	repo.db.Create(&interval)
	return interval
}

func (repo *Repository) StopInterval(stoppedAt string, id int64) model.Interval {
	var interval model.Interval
	repo.db.First(&interval, id)
	if interval.ID != 0 {
		interval.StoppedAt = stoppedAt
		repo.db.Save(&interval)
	}

	return interval
}

func (repo *Repository) GetIntervals() []model.Interval {
	var intervals []model.Interval
	repo.db.Find(&intervals)

	return intervals
}

func (repo *Repository) GetInterval(id int64) model.Interval {
	var interval model.Interval
	repo.db.First(&interval, id)

	return interval
}

func (repo *Repository) DeleteInterval(id int64) model.Interval {
	var interval model.Interval
	repo.db.First(&interval, id)
	if interval.ID != 0 {
		repo.db.Delete(&interval)
	}

	return interval
}
