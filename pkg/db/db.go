package db

import (
	"interval-action/pkg/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{Logger: logger.Default})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.Interval{})
	db.AutoMigrate(&model.Task{})

	return db
}
