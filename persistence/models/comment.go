package models

import (
	"devsmake/models"
	"time"

	"github.com/go-openapi/strfmt"
)

type Comments []Comment

type Comment struct {
	ID       uint64 `json:"id"`
	UUID     string `json:"uuid"`
	PostUUID string `json:"post_uuid"`
	AuthorID uint64 `json:"author_id"`
	Comment  string `json:"comment"`
	Deleted  bool   `json:"deleted"`
	Created  string `json:"created"`
}

func (c *Comment) PublicComment() *models.Comment {
	t, _ := time.Parse(time.RFC3339, c.Created)
	return &models.Comment{
		UUID:     c.UUID,
		AuthorID: int64(c.AuthorID),
		Comment:  &c.Comment,
		Created:  strfmt.DateTime(t),
	}
}

func (comments Comments) PublicComments() []*models.Comment {
	result := make([]*models.Comment, len(comments))
	for index, comment := range comments {
		result[index] = comment.PublicComment()
	}
	return result
}
