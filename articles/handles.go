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

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	articles := []Article{}
	db.Find(&articles)
	fmt.Println("Endpoint Hit: indexHandle")
	err = json.NewEncoder(w).Encode(articles)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateHandle(w http.ResponseWriter, r *http.Request) {
	db.AutoMigrate(&Article{})

	reqBody, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var article Article
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
