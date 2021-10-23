package main

import (
	"ConfigReader/models"
	"ConfigReader/reader"
	"fmt"
)

func main() {
	conf := models.Configuration{}
	reader := reader.New(&conf,
		reader.Settings{
			ConfigFile: "./application.conf.example",
			EnvPrefix:  "URMS",
		})
	err := reader.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conf)
}
