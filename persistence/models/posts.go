package models

import (
	"devsmake/models"

	"github.com/go-openapi/strfmt"
)

type Posts []Post

type Post struct {
	ID          uint64          `json:"id"`
	UUID        string          `json:"string"`
	AuthorID    uint64          `json:"author_id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Deleted     bool            `json:"deleted"`
	Modified    string          `json:"modified"`
	Created     strfmt.DateTime `json:"created"`
}

type Idea struct {
	ID          uint64          `json:"id"`
	UUID        string          `json:"string"`
	AuthorID    uint64          `json:"author_id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Likes       int             `json:"likes"`
	Dislikes    int             `json:"dislikes"`
	Submissions int             `json:"submissions"`
	Deleted     bool            `json:"deleted"`
	Modified    string          `json:"modified"`
	Created     strfmt.DateTime `json:"created"`
}

func (p *Post) PublicPost() *models.Idea {
	return &models.Idea{
		UUID:        p.UUID,
		AuthorID:    int64(p.AuthorID),
		Title:       &p.Title,
		Description: &p.Description,
		Likes:       0, // TODO
		Dislikes:    0,
		Submissions: 0,
		Created:     p.Created,
	}
}

func (posts Posts) PublicPosts() []*models.Idea {
	result := make([]*models.Idea, len(posts))
	for index, post := range posts {
		result[index] = post.PublicPost()
	}
	return result
}
