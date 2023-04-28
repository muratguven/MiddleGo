package databases

import (
	"log"
	"middle/pkg/config"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDatabase() *gorm.DB {
	config, er := config.NewConfig(".")
	if er != nil {
		log.Fatalln(er)
		panic("Failed to reading configuration file")
	}

	dbs, dbErr := gorm.Open(sqlserver.Open(string(config.ConnectionString)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})
	if dbErr != nil {

		log.Fatalln(dbErr)
		panic("Failed to connect to database")

	}

	return dbs
}
