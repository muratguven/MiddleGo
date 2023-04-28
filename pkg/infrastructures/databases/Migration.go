package databases

import (
	"middle/pkg/models"

	"gorm.io/gorm"
)

func MigrateDb(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{}, &models.Workflow{})
}
