package routes

import (
	"fmt"
	"interval-action/pkg/model"
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
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, Response{Data: nil, Status: http.StatusBadRequest, Error: &errorMessage})
		return
	}

	interval, err := r.repository.CreateInterval(body.StartAt, body.EndAt)
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusInternalServerError, Response{Data: nil, Status: http.StatusInternalServerError, Error: &errorMessage})
		return
	}

	intervalResponse := intervalToResponse(interval)

	c.JSON(http.StatusOK, Response{Data: intervalResponse, Status: http.StatusOK, Error: nil})
}

func (r *Route) StopInterval(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		StoppedAt string `json:"stopped_at"`
	}
	intervalID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, Response{Data: nil, Status: http.StatusBadRequest, Error: &errorMessage})
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, Response{Data: nil, Status: http.StatusBadRequest, Error: &errorMessage})
		return
	}

	interval, err := r.repository.StopInterval(body.StoppedAt, intervalID)
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusInternalServerError, Response{Data: nil, Status: http.StatusInternalServerError, Error: &errorMessage})
		return
	}

	intervalResponse := intervalToResponse(interval)

	c.JSON(http.StatusOK, Response{Data: intervalResponse, Status: http.StatusOK, Error: nil})
}

func (r *Route) GetIntervals(c *gin.Context) {
	intervals, err := r.repository.GetIntervals()
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusInternalServerError, Response{Data: nil, Status: http.StatusInternalServerError, Error: &errorMessage})
		return
	}

	intervalResponses := make([]model.IntervalResponse, len(intervals))
	for i, interval := range intervals {
		intervalResponses[i] = intervalToResponse(interval)
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, Response{Data: intervalResponses, Status: http.StatusOK, Error: nil})
}

func (r *Route) GetInterval(c *gin.Context) {
	intervalID := c.Param("id")
	fmt.Println(intervalID)
	intervalId, err := strconv.ParseInt(intervalID, 10, 64)
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, Response{Data: nil, Status: http.StatusBadRequest, Error: &errorMessage})
		return
	}

	interval, err := r.repository.GetInterval(uint(intervalId))
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusInternalServerError, Response{Data: nil, Status: http.StatusInternalServerError, Error: &errorMessage})
		return
	}

	intervalResponse := intervalToResponse(interval)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, Response{Data: intervalResponse, Status: http.StatusOK, Error: nil})
}

func intervalToResponse(interval model.Interval) model.IntervalResponse {
	return model.IntervalResponse{
		ID:        interval.ID,
		StartAt:   interval.StartAt,
		StoppedAt: interval.StoppedAt,
		EndAt:     interval.EndAt,
	}
}
