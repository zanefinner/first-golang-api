package accounts

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
)

var (
	t         *template.Template
	err       error
	config    map[string]string
	store     = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	session   *sessions.Session
	client    *http.Client
	remoteURL string
	db        *gorm.DB
)

//IndexHandle Resolves GET /accounts
func IndexHandle(w http.ResponseWriter, r *http.Request) {
	t, err = template.ParseFiles("templates/partials/head.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	config = map[string]string{
		"title": "Welcome",
	}
	err = t.Execute(w, config)
	if err != nil {
		fmt.Println(err)
		return
	}
	t, err = template.ParseFiles("templates/forms/signin.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	config = map[string]string{}
	err = t.Execute(w, config)
	if err != nil {
		fmt.Println(err)
		return
	}

	t, err = template.ParseFiles("templates/forms/signup.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	config = map[string]string{}
	err = t.Execute(w, config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Endpoint Hit: GET: /accounts")
}

//CreateHandle resolves POST /accounts/new
func CreateHandle(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", "zane:5245@/blog?charset=utf8&parseTime=True&loc=Local")
	Create(w, r)
	fmt.Println("Endpoint Hit: POST: /accounts/new")
}

//MatchHandle resolves POST /accounts/login
func MatchHandle(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", "zane:5245@/blog?charset=utf8&parseTime=True&loc=Local")
	Match(w, r)
	fmt.Println("Endpoint Hit: POST: /accounts/login")
}
