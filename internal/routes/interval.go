package routes

import (
	"fmt"
	"interval-action/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//	@Summary		Create a new interval
//	@Description	Create a new interval with the given start and end times
//	@Tags			Intervals
//	@Accept			json
//	@Produce		json
//	@Param			start_at	body		string					true	"Start time"	example:"2023-09-03T16:00:00Z"
//	@Param			end_at		body		string					true	"End time"		example:"2023-09-03T17:00:00Z"
//	@Success		200			{object}	model.IntervalResponse	"Successfully created interval"
//	@Router			/intervals [post]
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

//	@Summary		Stop an interval by ID
//	@Description	Stop an interval by its ID
//	@Tags			Intervals
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int						true	"Interval ID"
//	@Param			stopped_at	body		string					true	"Stop time"	example:"2023-09-03T16:30:00Z"
//	@Success		200			{object}	model.IntervalResponse	"Successfully stopped interval"
//	@Router			/intervals/{id} [put]
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

//	@Summary		Get all intervals
//	@Description	Get a list of all intervals
//	@Tags			Intervals
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	model.IntervalResponse	"List of intervals"
//	@Router			/intervals [get]
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

//	@Summary		Get an interval by ID
//	@Description	Get an interval by its ID
//	@Tags			Intervals
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int						true	"Interval ID"
//	@Success		200	{object}	model.IntervalResponse	"Interval details"
//	@Router			/intervals/{id} [get]
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
