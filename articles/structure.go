package articles

type Article struct {
	ID      int    `json:"id"`
	User    string `json:"user"`
	Content string `json:"content"`
}
