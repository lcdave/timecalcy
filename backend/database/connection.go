package database

import (
	"log"
	"timecalcy/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://admin:asdf@localhost:5432/timecalcy-db"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&models.Product{})

	if err != nil {
		return nil
	}

	product := models.Product{
		Name: "test",
	}

	db.Create(&product)

	return db
}
