package databases

import (
	"middle/pkg/domains"

	"gorm.io/gorm"
)

func MigrateDb(db *gorm.DB) error {
	return db.AutoMigrate(&domains.User{}, &domains.Workflow{})
}
