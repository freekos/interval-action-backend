package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Route) CreateTask(c *gin.Context) {
	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		IntervalId  int64  `json:"interval_id"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := r.repository.CreateTask(body.Title, body.Description, body.IntervalId)

	c.JSON(http.StatusOK, task)
}

func (r *Route) GetTasks(c *gin.Context) {
	intervalID := c.DefaultQuery("id", "0")
	tasks := r.repository.GetTasks(intervalID)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, tasks)
}

func (r *Route) GetTask(c *gin.Context) {
	taskID := c.DefaultQuery("id", "0")
	task := r.repository.GetTask(taskID)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, task)
}

func (r *Route) DeleteTask(c *gin.Context) {
	taskID := c.DefaultQuery("id", "0")
	task := r.repository.DeleteTask(taskID)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, task)
}
