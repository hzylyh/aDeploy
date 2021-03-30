package main

import (
	"aDeploy/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func main() {
	generateTable()
	router := MapRoutes()
	router.Run(":8089")
}

func generateTable() {
	var err error
	DB, err = gorm.Open(sqlite.Open("aDeploy.db"), &gorm.Config{})
	if err = DB.AutoMigrate(models.Models...); nil != err {
		log.Fatal("auto migrate tables failed: " + err.Error())
	}
}
