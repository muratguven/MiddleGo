package databases

import (
	"log"
	"middle/pkg/models"
	"middle/pkg/utils"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	config, er := utils.LoadConfig(".")
	if er != nil {
		log.Fatalln(er)
		panic("Failed to reading configuration file")
	}

	dbs, dbErr := gorm.Open(sqlserver.Open(string(config.ConnectionString)), &gorm.Config{})
	if dbErr != nil {

		log.Fatalln(dbErr)
		panic("Failed to connect to database")

	}

	dbs.AutoMigrate(&models.User{})

	DB = dbs
}
