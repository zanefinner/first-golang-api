package articles

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error
var reqBody []byte

//IndexHandle returns all records
func IndexHandle(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", "zane:5245@/blog?charset=utf8&parseTime=True&loc=Local")
	Read(w, r)

}

//CreateHandle creates a record
func CreateHandle(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("mysql", "zane:5245@/blog?charset=utf8&parseTime=True&loc=Local")
	Create(w, r)
}
