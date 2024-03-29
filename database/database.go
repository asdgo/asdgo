package database

import (
	"gorm.io/gorm"
)

var Instance *Database

type Database struct {
	*gorm.DB
}

func New(dialector gorm.Dialector) {
	db, err := gorm.Open(dialector)

	if err != nil {
		panic(err)
	}

	Instance = &Database{
		db,
	}
}
