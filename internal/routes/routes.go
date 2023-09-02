package routes

import (
	"interval-action/pkg/repository"

	"github.com/gin-gonic/gin"
)

type Route struct {
	repository *repository.Repository
}

type Response struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
	Error  *string     `json:"error"`
}

func Init(repository *repository.Repository) *gin.Engine {
	r := gin.Default()

	route := Route{repository}

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/intervals", route.CreateInterval)
			v1.POST("/intervals/stop", route.StopInterval)
			v1.GET("/intervals", route.GetIntervals)
			v1.GET("/intervals/:id", route.GetInterval)

			v1.POST("/tasks", route.CreateTask)
			v1.DELETE("/tasks/:id", route.DeleteTask)
			v1.GET("/tasks", route.GetTasks)
			v1.GET("/tasks/:id", route.GetTask)
		}
	}

	return r
}
