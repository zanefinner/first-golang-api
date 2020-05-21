package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zanefinner/golang-blog/accounts"
	"github.com/zanefinner/golang-blog/articles"
	"github.com/zanefinner/golang-blog/base"
)

//Start makes all the routes
func Start() {
	myRouter := mux.NewRouter().StrictSlash(true)
	//Base Routes
	myRouter.HandleFunc("/", base.IndexHandle).Methods("GET")

	//Article Routes
	myRouter.HandleFunc("/articles", articles.IndexHandle).Methods("GET")
	myRouter.HandleFunc("/articles", articles.CreateHandle).Methods("POST")

	//Account Routes
	myRouter.HandleFunc("/accounts", accounts.IndexHandle).Methods("GET")
	myRouter.HandleFunc("/accounts/new", accounts.CreateHandle).Methods("POST")
	myRouter.HandleFunc("/accounts/login", accounts.MatchHandle).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
