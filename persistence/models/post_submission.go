package models

import (
	"devsmake/models"
	"time"

	"github.com/go-openapi/strfmt"
)

type IdeaSubmissions []IdeaSubmission

type IdeaSubmission struct {
	ID       uint64 `json:"id"`
	UUID     string `json:"uuid"`
	PostUUID string `json:"post_uuid"`
	AuthorID uint64 `json:"author_id"`
	Comment  string `json:"comment"`
	Rating   int    `json:"rating"`
	Deleted  bool   `json:"deleted"`
	Modified string `json:"modified"`
	Created  string `json:"created"`
}

func (is *IdeaSubmission) PublicIdeaSubmission() *models.Submission {
	t, _ := time.Parse(time.RFC3339, is.Created)
	return &models.Submission{
		UUID:     is.UUID,
		AuthorID: int64(is.AuthorID),
		Comment:  &is.Comment,
		Rating:   int64(is.Rating),
		Created:  strfmt.DateTime(t),
	}
}

func (idea_submissions IdeaSubmissions) PublicIdeaSubmissions() []*models.Submission {
	result := make([]*models.Submission, len(idea_submissions))
	for index, idea_submission := range idea_submissions {
		result[index] = idea_submission.PublicIdeaSubmission()
	}
	return result
}
