package articles

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error
var reqBody []byte

//IndexHandle resolves GET /articles
func IndexHandle(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", "zane:5245@/blog?charset=utf8&parseTime=True&loc=Local")
	Read(w, r)
	fmt.Println("Endpoint Hit: indexHandle")
}

//CreateHandle resolves POST /articles
func CreateHandle(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", "zane:5245@/blog?charset=utf8&parseTime=True&loc=Local")
	Create(w, r)
	fmt.Println("Endpoint Hit: Creating New Post")
}
