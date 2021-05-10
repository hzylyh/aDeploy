/*
 * @Description:
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-04-13 19:08:12
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-10 17:34:16
 */
package main

import (
	"aDeploy/conf"
	"aDeploy/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	generateTable()
	router := MapRoutes()
	router.Run(":8089")
}

func generateTable() {
	var err error
	conf.DB, err = gorm.Open(sqlite.Open("aDeploy.db"), &gorm.Config{})
	if err = conf.DB.AutoMigrate(models.Models...); nil != err {
		log.Fatal("auto migrate tables failed: " + err.Error())
	}
}
