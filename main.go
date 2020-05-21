package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zanefinner/first-golang-api/router"
)

var db *gorm.DB
var err error
var reqBody []byte

func main() {
	if err != nil {
		log.Println("Connection Failed to Open")
		return
	}

	log.Println("Connection Established")
	log.Println("Server Start")

	router.Start()
}
