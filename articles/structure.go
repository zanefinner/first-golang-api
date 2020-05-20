package articles

//Article is the json format of records
type Article struct {
	ID      int    `json:"id"`
	User    string `json:"user"`
	Content string `json:"content"`
}
