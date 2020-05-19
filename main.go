package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error
var reqBody []byte

type article struct {
    ID      int    `json:"id"`
    User    string `json:"user"`
    Content string `json:"content"`
}


func main() {

    db, err = gorm.Open("mysql", "zane:5245@/blog?charset=utf8&parseTime=True&loc=Local")

    if err != nil {
        log.Println("Connection Failed to Open")
        return
    }

    log.Println("Connection Established")
    db.AutoMigrate(&article{})

    log.Println("Server Start")

    myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/", homeHandle).Methods("GET")
    myRouter.HandleFunc("/articles", indexHandle).Methods("GET")
    myRouter.HandleFunc("/articles", createHandle).Methods("POST")

    log.Fatal(http.ListenAndServe(":10000", myRouter))
}