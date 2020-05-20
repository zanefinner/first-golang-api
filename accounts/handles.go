package accounts

import (
	"fmt"
	"net/http"
)

//IndexHandle Resolves GET /accounts
func IndexHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Log in/Sign up form will be here"))
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
	// Does the username or pass match a record?
	// If yes, get password from the record. Does that match?
	//  If yes, set logged in session
	// If not, return error and prompt to try again
	//If not, state that the user doesn't exist
}
