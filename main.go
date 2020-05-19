package main

import (
    "log"
	"net/http"
	"github.com/zanefinner/first-golang-api/articles"
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error
var reqBody []byte


func main() {

    db, err = gorm.Open("mysql", "zane:5245@/blog?charset=utf8&parseTime=True&loc=Local")

    if err != nil {
        log.Println("Connection Failed to Open")
        return
    }

    log.Println("Connection Established")
    log.Println("Server Start")

    myRouter := mux.NewRouter().StrictSlash(true)

    //myRouter.HandleFunc("/", articles.homeHandle).Methods("GET")
    myRouter.HandleFunc("/articles", articles.IndexHandle).Methods("GET")
    myRouter.HandleFunc("/articles", articles.CreateHandle).Methods("POST")

    log.Fatal(http.ListenAndServe(":10000", myRouter))
}