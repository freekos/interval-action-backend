package routes

import (
	"interval-action/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Route) CreateTask(c *gin.Context) {
	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		IntervalID  string `json:"interval_id"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, Response{Data: nil, Status: http.StatusBadRequest, Error: &errorMessage})
		return
	}

	intervalId, err := strconv.ParseUint(body.IntervalID, 10, 64)
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, Response{Data: nil, Status: http.StatusBadRequest, Error: &errorMessage})
		return
	}

	task, err := r.repository.CreateTask(body.Title, body.Description, uint(intervalId))
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusInternalServerError, Response{Data: nil, Status: http.StatusInternalServerError, Error: &errorMessage})
		return
	}

	taskResponse := taskToResponse(task)

	c.JSON(http.StatusOK, Response{Data: taskResponse, Status: http.StatusOK, Error: nil})
}

func (r *Route) GetTasks(c *gin.Context) {
	var body struct {
		IntervalID string `json:"interval_id"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, Response{Data: nil, Status: http.StatusBadRequest, Error: &errorMessage})
		return
	}
	intervalId, err := strconv.ParseUint(body.IntervalID, 10, 64)
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, Response{Data: nil, Status: http.StatusBadRequest, Error: &errorMessage})
		return
	}

	tasks, err := r.repository.GetTasks(uint(intervalId))
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusInternalServerError, Response{Data: nil, Status: http.StatusInternalServerError, Error: &errorMessage})
		return
	}

	var tasksResponse = make([]model.TaskResponse, len(tasks))
	for i, task := range tasks {
		tasksResponse[i] = taskToResponse(task)
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, Response{Data: tasksResponse, Status: http.StatusOK, Error: nil})
}

func (r *Route) GetTask(c *gin.Context) {
	taskID := c.Param("id")
	taskId, err := strconv.ParseUint(taskID, 10, 64)
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, Response{Data: nil, Status: http.StatusBadRequest, Error: &errorMessage})
		return
	}

	task, err := r.repository.GetTask(uint(taskId))
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusInternalServerError, Response{Data: nil, Status: http.StatusInternalServerError, Error: &errorMessage})
		return
	}

	taskResponse := taskToResponse(task)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, Response{Data: taskResponse, Status: http.StatusOK, Error: nil})
}

func (r *Route) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	taskId, err := strconv.ParseUint(taskID, 10, 64)
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, Response{Data: nil, Status: http.StatusBadRequest, Error: &errorMessage})
		return
	}

	task, err := r.repository.DeleteTask(uint(taskId))
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusInternalServerError, Response{Data: nil, Status: http.StatusInternalServerError, Error: &errorMessage})
		return
	}

	taskResponse := taskToResponse(task)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, Response{Data: taskResponse, Status: http.StatusOK, Error: nil})
}

func taskToResponse(task model.Task) model.TaskResponse {
	return model.TaskResponse{
		ID:          task.ID,
		IntervalID:  task.IntervalID,
		Title:       task.Title,
		Description: task.Description,
	}
}
