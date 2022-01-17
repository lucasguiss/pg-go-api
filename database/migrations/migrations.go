package migrations

import (
	"github.com/lucasguiss/pg-go-api/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		models.Book{},
	)
}
