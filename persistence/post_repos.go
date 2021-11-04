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

func (r *PostRepos) GetPostSubmissions(uuid string) (models.Submissions, error) {
	var submissions = models.Submissions{}

	query := `
		SELECT id, uuid, post_uuid, author_id, comment, deleted, modified, created 
		FROM posts_submissions WHERE deleted = 0 AND post_uuid = ?
  	`

	results, err := r.db.Query(query, uuid)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var submission models.Submission

		results.Scan(
			&submission.ID,
			&submission.UUID,
			&submission.PostUUID,
			&submission.AuthorID,
			&submission.Comment,
			&submission.Deleted,
			&submission.Modified,
			&submission.Created,
		)

		submissions = append(submissions, submission)
	}
	return submissions, nil
}

func (r *PostRepos) GetPostRatings(uuid string, rating interfaces.PostRating) (*int, error) {
	var count int

	err := r.db.QueryRow("SELECT COUNT(*) FROM posts_ratings WHERE uuid = ? AND rating = ?", uuid, rating).Scan(&count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

func (r *PostRepos) GetIdea(uuid string) (*models.Idea, error) {
	post, err := r.GetPost(uuid)
	if err != nil {
		return nil, err
	}

	likes, err := r.GetPostRatings(uuid, interfaces.Like)
	if err != nil {
		return nil, err
	}

	dislikes, err := r.GetPostRatings(uuid, interfaces.Dislike)
	if err != nil {
		return nil, err
	}

	submissions, err := r.GetPostSubmissions(uuid)
	if err != nil {
		return nil, err
	}

	return &models.Idea{
		ID:          post.ID,
		UUID:        post.UUID,
		AuthorID:    post.AuthorID,
		Title:       post.Title,
		Description: post.Description,
		Likes:       *likes,
		Dislikes:    *dislikes,
		Submissions: len(submissions),
		Deleted:     post.Deleted,
		Modified:    post.Modified,
		Created:     post.Created,
	}, nil
}
