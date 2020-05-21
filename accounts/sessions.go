package accounts

import (
	"fmt"
	"net/http"
)

//Auth starts a session for when you log in
func Auth(iud int) {
	session, err := store.Get(r, "auth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set some session values.
	session.Values["iud"] = int
	// Save it before we write to the response/return from the handler.
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Auth session started")
}
