package models

type Posts []Post

type Post struct {
	ID          uint64 `json:"id"`
	UUID        string `json:"uuid"`
	AuthorID    uint64 `json:"author_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Deleted     bool   `json:"deleted"`
	Modified    string `json:"modified"`
	Created     string `json:"created"`
}
