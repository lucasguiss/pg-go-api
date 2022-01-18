package database

import (
	"log"
	"time"

	"github.com/lucasguiss/pg-go-api/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	str := "host=localhost port=5438 user=postgres password=postgres sslmode=disable dbname=books"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}
