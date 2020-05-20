package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zanefinner/first-golang-api/articles"
)

var db *gorm.DB
var err error
var reqBody []byte

func main() {
	if err != nil {
		log.Println("Connection Failed to Open")
		return
	}

	//db.AutoMigrate(&Article{})

	log.Println("Connection Established")
	log.Println("Server Start")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/articles", articles.IndexHandle).Methods("GET")
	myRouter.HandleFunc("/articles", articles.CreateHandle).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
