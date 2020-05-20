package base

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	t      *template.Template
	err    error
	config map[string]string
)

//IndexHandle Resolves the site's root
func IndexHandle(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Website Root"))
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
	t, err = template.ParseFiles("templates/base/landing_page.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	config = map[string]string{
		"name": "Zane",
	}
	err = t.Execute(w, config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Endpoint Hit: Website Root")
}
