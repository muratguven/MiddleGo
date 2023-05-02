package api

import (
	"middle/pkg/config"
	"middle/pkg/infrastructures/databases"
	"middle/pkg/infrastructures/logs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

type Application struct {
	Db     *gorm.DB
	config *config.Config
}

// Application initializer func
func Run(cfg *config.Config) {

	app := new(Application)
	app.config = cfg
	app.Db = databases.ConnectDatabase()
	databases.MigrateDb(app.Db)

	// WebServer start

	server := fiber.New()
	server.Use(limiter.New(limiter.Config{Max: 100}), cors.New(cors.ConfigDefault), logger.New(logs.CustomApiLogger()))

	app.Routes(server)
	server.Listen(":" + app.config.Port)

}
