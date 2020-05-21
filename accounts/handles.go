package accounts

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var (
	t         *template.Template
	err       error
	config    map[string]string
	store     = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	session   *sessions.Session
	client    *http.Client
	remoteURL string
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
	t, err = template.ParseFiles("templates/forms/signin.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	config = map[string]string{}
	err = t.Execute(w, config)

	fmt.Println("Endpoint Hit: GET: /accounts")
}

//CreateHandle resolves POST /accounts/new
func CreateHandle(w http.ResponseWriter, r *http.Request) {
	// Is the username, password, and email all valid?
	// If so, make account and prompt to log in
	// If not, redirect back to page with error
	fmt.Println("Endpoint Hit: POST: /accounts/new")
}

//MatchHandle resolves POST /accounts/login
func MatchHandle(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	log.Println(r.Form)
	fmt.Println("Endpoint Hit: POST: /accounts/login")
}
