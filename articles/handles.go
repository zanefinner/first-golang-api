package articles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error
var reqBody []byte

type article struct {
	ID      int    `json:"id"`
	User    string `json:"user"`
	Content string `json:"content"`
}

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
