package base

import (
	"fmt"
	"html/template"
	"net/http"
)

//IndexHandle Resolves the site's root
func IndexHandle(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Website Root"))
	t, err := template.ParseFiles("templates/landing_page.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	config := map[string]string{
		"name": "Zane",
	}
	err = t.Execute(w, config)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Endpoint Hit: Website Root")
}
