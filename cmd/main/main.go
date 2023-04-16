package main

import (
	"fmt"
	"log"
	"middle/pkg/utils"

	"gopkg.in/ini.v1"
)

func main() {

	config, err := ini.Load("appdev.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println(config.Section("").Key("application_mode").String())
	fmt.Println(config.Section("").Key("default_cnn").String())

	fmt.Println("Advanced go project")

	//using Viper config (app.env file)

	c, e := utils.LoadConfig(".")
	if e != nil {
		log.Fatal("Can not load config", e)
	}
	fmt.Println(c.ConnectionString)
}
