package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Route) CreateInterval(c *gin.Context) {
	var body struct {
		StartAt string `json:"start_at"`
		EndAt   string `json:"end_at"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	interval := r.repository.CreateInterval(body.StartAt, body.EndAt)

	c.JSON(http.StatusOK, interval)
}

func (r *Route) StopInterval(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	intervalID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval ID"})
		return
	}

	var body struct {
		StoppedAt string `json:"stopped_at"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	interval := r.repository.StopInterval(body.StoppedAt, intervalID)

	c.JSON(http.StatusOK, interval)
}

func (r *Route) GetIntervals(c *gin.Context) {
	intervals := r.repository.GetIntervals()

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, intervals)
}

func (r *Route) GetInterval(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	intervalID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval ID"})
		return
	}

	interval := r.repository.GetInterval(intervalID)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, interval)
}
