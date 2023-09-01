package repository

import (
	"interval-action/pkg/model"
)

func (repo *Repository) CreateTask(title string, description string, intervalId int64) model.Task {
	task := model.Task{
		Title:       title,
		Description: description,
	}

	repo.db.Create(&task)

	var interval model.Interval
	if repo.db.First(&interval, intervalId).Error == nil {
		interval.Tasks = append(interval.Tasks, task)
		repo.db.Save(&interval)
	}

	return task
}

func (repo *Repository) GetTasks(intervalId string) []model.Task {
	var tasks []model.Task
	var interval model.Interval
	if repo.db.First(&interval, intervalId).Error == nil {
		repo.db.Find(&tasks, interval.Tasks)
	}

	return tasks
}

func (repo *Repository) GetTask(id string) model.Task {
	var task model.Task
	repo.db.First(&task, id)

	return task
}

func (repo *Repository) DeleteTask(id string) model.Task {
	var task model.Task
	repo.db.First(&task, id)
	if task.ID != 0 {
		repo.db.Delete(&task)
	} else {
	}

	return task
}
