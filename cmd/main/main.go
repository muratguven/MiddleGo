package main

import (
	"fmt"
	"log"
	"middle/pkg/api"
	"middle/pkg/config"
	"middle/pkg/infrastructures/logs"
	"os"
)

func init() {
	fmt.Println("Application starting...")
}

func main() {
	fmt.Println("Application started!")

	logs.Filelogger().Info().Msg("Test info log")

	//using Viper config (app.env file)

	cfg, err := config.NewConfig("../../")
	if err != nil {
		log.Fatal("Can not load config", err)
		os.Exit(1)
	}
	api.Run(cfg)

}
