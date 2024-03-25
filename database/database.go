package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func New() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		panic(err)
	}

	Instance = db
}
