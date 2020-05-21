package accounts

import (
	"fmt"
	"net/http"
)

//Create adds a record to the db
func Create(w http.ResponseWriter, r *http.Request) {
	var account Account
	err = r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	account.Email = r.Form["email"][0]
	account.Username = r.Form["username"][0]
	account.Password = r.Form["password"][0]
	db.Create(&account)
}

//Match compares a record in the db with a form entry
func Match(w http.ResponseWriter, r *http.Request) {
	var account Account
	err = r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	//db.AutoMigrate(&Account{})
	db.Where("username = ? AND password = ?", r.Form["username"], r.Form["password"]).Find(&account)
	if account.ID == 0 {
		fmt.Println("Failed to log in")
		//retry (return login page with error)
	} else {
		fmt.Println("Logged in")
		Auth(account.ID)
	}
}
