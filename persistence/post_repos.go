package persistence

import (
	"database/sql"
	"devsmake/persistence/interfaces"
	"devsmake/persistence/models"

	uuid "github.com/satori/go.uuid"
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

func (r *PostRepos) GetPosts(limit uint64, offset uint64, search string) (models.Posts, error) {
	var posts = models.Posts{}

	query := `
		SELECT id, uuid, author_id, title, description, deleted, modified, created 
		FROM posts WHERE deleted = 0 AND (title LIKE ? OR description LIKE ?) 
		ORDER BY id DESC LIMIT ?, ?
  	`

	search = "%" + search + "%"
	results, err := r.db.Query(query, search, search, offset, limit)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var post models.Post

		results.Scan(
			&post.ID,
			&post.UUID,
			&post.AuthorID,
			&post.Title,
			&post.Description,
			&post.Deleted,
			&post.Modified,
			&post.Created,
		)

		posts = append(posts, post)
	}
	return posts, nil

}

func (r *PostRepos) CreatePost(authorId uint64, title string, desc string) (string, error) {
	uuid := uuid.NewV4().String()
	query := `
		INSERT INTO posts (uuid, author_id, title, description) 
		VALUES (?, ?, ?, ?)
	`
	_, err := r.db.Query(query, uuid, authorId, title, desc)

	return uuid, err
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

	err := r.db.QueryRow("SELECT COUNT(*) FROM posts_ratings WHERE post_uuid = ? AND rating = ?", uuid, rating).Scan(&count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

func (r *PostRepos) RatePost(raterID uint64, postUUID string, rating interfaces.PostRating) error {
	var count int

	err := r.db.QueryRow("SELECT COUNT(*) FROM posts_ratings WHERE post_uuid = ? AND rater_id = ?", postUUID, raterID).Scan(&count)
	if err != nil {
		return err
	}

	if count != 0 {
		_, err := r.db.
			Query("UPDATE posts_ratings SET rating = ? WHERE post_uuid = ? AND rater_id = ?",
				rating, postUUID, raterID)

		return err
	}

	_, err = r.db.
		Query("INSERT INTO posts_ratings (post_uuid, rater_id, rating) VALUES (?, ?, ?)", postUUID, raterID, rating)

	return err
}

func (r *PostRepos) GetIdea(uuid string) (*models.Idea, error) {
	// TODO make more performant

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

func (r *PostRepos) GetIdeas(limit uint64, offset uint64, query string) (models.Ideas, error) {
	var ideas = models.Ideas{}

	posts, err := r.GetPosts(limit, offset, query)
	if err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return ideas, nil
	}

	// TODO make more performant
	for _, post := range posts {
		likes, _ := r.GetPostRatings(post.UUID, interfaces.Like)
		dislikes, _ := r.GetPostRatings(post.UUID, interfaces.Dislike)
		submissions, _ := r.GetPostSubmissions(post.UUID)

		ideas = append(ideas,
			models.Idea{
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
			})
	}
	return ideas, nil
}
