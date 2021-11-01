package persistence

import (
	"database/sql"
	"devsmake/persistence/interfaces"
	"devsmake/persistence/models"
)

type PostRepos struct {
	db *sql.DB
}

func NewPostRepositories(db *sql.DB) *PostRepos {
	return &PostRepos{db}
}

var _ interfaces.PostRepositories = &PostRepos{}

func (r *PostRepos) GetPost(uuid string) (*models.Post, error) {
	var post models.Post

	query := `
		SELECT id, uuid, author_id, title, description, deleted, modified, created
		FROM posts WHERE deleted = 0 AND uuid = ?
  	`

	err := r.db.
		QueryRow(query, uuid).
		Scan(
			&post.ID,
			&post.UUID,
			&post.AuthorID,
			&post.Title,
			&post.Description,
			&post.Deleted,
			&post.Modified,
			&post.Created,
		)

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *PostRepos) GetPostRatings(uuid string, rating interfaces.PostRating) (*int, error) {
	var count int

	err := r.db.QueryRow("SELECT COUNT(*) FROM posts_ratings WHERE uuid = ? AND rating = ?", uuid, rating).Scan(&count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}
