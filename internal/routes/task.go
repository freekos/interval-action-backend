package routes

import (
	"interval-action/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//	@Summary		Create a new task
//	@Description	Create a new task with the provided title, description, and interval ID
//	@Tags			Tasks
//	@Accept			json
//	@Produce		json
//	@Param			title		body		string	true	"Task title"
//	@Param			description	body		string	true	"Task description"
//	@Param			interval_id	body		string	true	"Interval ID"	example:"1"
//	@Success		200			{object}	model.TaskResponse
//	@Router			/tasks [post]
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

//	@Summary		Get tasks by interval ID
//	@Description	Get a list of tasks by the provided interval ID
//	@Tags			Tasks
//	@Accept			json
//	@Produce		json
//	@Param			interval_id	query	int	true	"Interval ID"	example:"1"
//	@Success		200			{array}	model.TaskResponse
//	@Router			/tasks [get]
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

//	@Summary		Get a task by ID
//	@Description	Get a task by its ID
//	@Tags			Tasks
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Task ID"
//	@Success		200	{object}	model.TaskResponse
//	@Router			/tasks/{id} [get]
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

//	@Summary		Delete a task by ID
//	@Description	Delete a task by its ID
//	@Tags			Tasks
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Task ID"
//	@Success		200	{object}	model.TaskResponse
//	@Router			/tasks/{id} [delete]
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
