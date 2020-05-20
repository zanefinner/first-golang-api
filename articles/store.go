package articles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Read returns records
func Read(w http.ResponseWriter, r *http.Request) {
	articles := []Article{}
	db.Find(&articles)
	fmt.Println("Endpoint Hit: indexHandle")
	err = json.NewEncoder(w).Encode(articles)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//Create adds to the db
func Create(w http.ResponseWriter, r *http.Request) {
	reqBody, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var article Article
	err = json.Unmarshal(reqBody, &article)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&article)

	fmt.Println("Endpoint Hit: Creating New Post")
	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		fmt.Println(err)
		return
	}
}
