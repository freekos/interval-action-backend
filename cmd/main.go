package main

import (
	"fmt"
	"interval-action/internal/routes"
	"interval-action/pkg/db"
	"interval-action/pkg/repository"
	"net/http"
	"os"

	_ "interval-action/docs"
)

//	@title		Interval-Action API
//	@version	0.0.20
//	@host		localhost:8080
//	@BasePath	/api/v1

func main() {
	db := db.Init(os.Getenv("DB_URL"))

	repository := repository.New(db)

	r := routes.Init(&repository)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
