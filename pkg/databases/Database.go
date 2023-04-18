package databases

import (
	"log"
	"middle/pkg/utils"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	config, er := utils.LoadConfig(".")
	if er != nil {
		panic("Failed to reading configuration file")
		log.Fatalln(er)
	}

	dbs, dbErr := gorm.Open(sqlserver.Open(string(config.ConnectionString)), &gorm.Config{})
	if dbErr != nil {
		panic("Failed to connect to database")
		log.Fatalln(dbErr)
	}
	

	DB = dbs
}
