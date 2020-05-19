package article

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	articles := []article{}
	db.Find(&articles)
	fmt.Println("Endpoint Hit: indexHandle")
	err = json.NewEncoder(w).Encode(articles)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateHandle(w http.ResponseWriter, r *http.Request) {
	reqBody, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var article article
	err = json.Unmarshal(reqBody, &article)
	if err != nil {
		log.Fatal(err)
	}
	db.Create(&article)
	fmt.Println("Endpoint Hit: Creating New Post")
	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		log.Fatal(err)
	}
}
