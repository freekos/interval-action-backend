package repository

import (
	"interval-action/pkg/model"
)

func (repo *Repository) CreateTask(title string, description string, intervalId uint) (model.Task, error) {
	var interval model.Interval
	if err := repo.db.First(&interval, intervalId).Error; err != nil {
		return model.Task{}, err
	}

	task := model.Task{
		Title:       title,
		Description: description,
		IntervalID:  intervalId,
	}

	if err := repo.db.Create(&task).Error; err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (repo *Repository) GetTasks(intervalId uint) ([]model.Task, error) {
	var tasks []model.Task

	var interval model.Interval
	if err := repo.db.First(&interval, intervalId).Error; err != nil {
		return nil, err
	}

	if err := repo.db.Where("interval_id = ?", interval.ID).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (repo *Repository) GetTask(id uint) (model.Task, error) {
	var task model.Task
	if err := repo.db.First(&task, id).Error; err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (repo *Repository) DeleteTask(id uint) (model.Task, error) {
	var task model.Task
	if err := repo.db.First(&task, id).Error; err != nil {
		return model.Task{}, err
	}

	repo.db.Delete(&task)

	return task, nil
}
