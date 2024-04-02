package database

import "gorm.io/gorm"

var Instance *Database

type Database struct {
	*gorm.DB
}

func New(db *gorm.DB) {
	Instance = &Database{db}
}
