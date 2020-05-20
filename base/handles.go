package base

import (
	"net/http"
	"fmt"
)
//IndexHandle Resolves the site's root
func IndexHandle(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Website Root"))
	fmt.Println("Endpoint Hit: Website Root")
}