package repository

import (
	"interval-action/pkg/model"
)

func (repo *Repository) CreateInterval(startAt string, endAt string) (model.Interval, error) {
	interval := model.Interval{StartAt: startAt, EndAt: endAt}

	if err := repo.db.Create(&interval).Error; err != nil {
		return model.Interval{}, err
	}

	return interval, nil
}

func (repo *Repository) StopInterval(stoppedAt string, id int64) (model.Interval, error) {
	var interval model.Interval
	if err := repo.db.First(&interval, id).Error; err != nil {
		return model.Interval{}, err
	}

	interval.StoppedAt = stoppedAt
	if err := repo.db.Save(&interval).Error; err != nil {
		return model.Interval{}, err
	}

	return interval, nil
}

func (repo *Repository) GetIntervals() ([]model.Interval, error) {
	var intervals []model.Interval
	if err := repo.db.Find(&intervals).Error; err != nil {
		return nil, err
	}

	return intervals, nil
}

func (repo *Repository) GetInterval(id uint) (model.Interval, error) {
	var interval model.Interval
	if err := repo.db.First(&interval, id).Error; err != nil {
		return model.Interval{}, err
	}

	return interval, nil
}

func (repo *Repository) DeleteInterval(id int64) (model.Interval, error) {
	var interval model.Interval
	if err := repo.db.First(&interval, id).Error; err != nil {
		return model.Interval{}, err
	}

	repo.db.Delete(&interval)

	return interval, nil
}
