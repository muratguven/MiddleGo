package main

import (
	"fmt"
	"log"
	"middle/pkg/databases"
	"middle/pkg/utils"
)

func init() {
	fmt.Println("Application starting...")
}

func main() {
	fmt.Println("Application started!")

	//using Viper config (app.env file)

	c, e := utils.LoadConfig(".")
	if e != nil {
		log.Fatal("Can not load config", e)
	}
	fmt.Println(c.ConnectionString)

	databases.ConnectDatabase()
}
