package models

import (
	"devsmake/models"
	"time"

	"github.com/go-openapi/strfmt"
)

type Ideas []Idea

type Idea struct {
	ID          uint64 `json:"id"`
	UUID        string `json:"uuid"`
	AuthorID    uint64 `json:"author_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
	Submissions int    `json:"submissions"`
	Deleted     bool   `json:"deleted"`
	Modified    string `json:"modified"`
	Created     string `json:"created"`
}

func (i *Idea) PublicIdea() *models.Idea {
	t, _ := time.Parse(time.RFC3339, i.Created)
	return &models.Idea{
		UUID:        i.UUID,
		AuthorID:    int64(i.AuthorID),
		Title:       &i.Title,
		Description: &i.Description,
		Likes:       int64(i.Likes),
		Dislikes:    int64(i.Dislikes),
		Submissions: int64(i.Submissions),
		Created:     strfmt.DateTime(t),
	}
}

func (ideas Ideas) PublicIdeas() []*models.Idea {
	result := make([]*models.Idea, len(ideas))
	for index, idea := range ideas {
		result[index] = idea.PublicIdea()
	}
	return result
}
