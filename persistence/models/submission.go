package models

type Submissions []Submission

type Submission struct {
	ID       uint64 `json:"id"`
	UUID     string `json:"uuid"`
	PostUUID string `json:"post_uuid"`
	AuthorID uint64 `json:"author_id"`
	Comment  string `json:"comment"`
	Deleted  bool   `json:"deleted"`
	Modified string `json:"modified"`
	Created  string `json:"created"`
}
