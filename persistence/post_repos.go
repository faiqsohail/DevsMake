package persistence

import (
	"database/sql"
	"devsmake/persistence/interfaces"
	"devsmake/persistence/models"
	"math"

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

func (r *PostRepos) CreateSubmission(authorId uint64, postUUID string, comment string) (string, error) {
	uuid := uuid.NewV4().String()
	query := `
		INSERT INTO posts_submissions (uuid, author_id, post_uuid, comment) 
		VALUES (?, ?, ?, ?)
	`
	_, err := r.db.Query(query, uuid, authorId, postUUID, comment)

	return uuid, err
}

func (r *PostRepos) GetSubmission(uuid string) (*models.Submission, error) {
	var submission models.Submission

	query := `
		SELECT id, uuid, post_uuid, author_id, comment, deleted, modified, created 
		FROM posts_submissions WHERE deleted = 0 AND uuid = ?
  	`

	err := r.db.
		QueryRow(query, uuid).
		Scan(
			&submission.ID,
			&submission.UUID,
			&submission.PostUUID,
			&submission.AuthorID,
			&submission.Comment,
			&submission.Deleted,
			&submission.Modified,
			&submission.Created,
		)

	if err != nil {
		return nil, err
	}

	return &submission, nil
}

func (r *PostRepos) GetSubmissions(postUUID string) (models.Submissions, error) {
	var submissions = models.Submissions{}

	query := `
		SELECT id, uuid, post_uuid, author_id, comment, deleted, modified, created 
		FROM posts_submissions WHERE deleted = 0 AND post_uuid = ?
  	`

	results, err := r.db.Query(query, postUUID)
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

func (r *PostRepos) GetSubmissionRating(postSubmissionUUID string) (*int, error) {
	var total_ratings int
	var sum_rating int

	query := `
		SELECT COUNT(*) as total_ratings, SUM(rating) as sum_rating 
		FROM posts_submissions_ratings 
		WHERE post_submission_uuid = ? AND rating > 0
  	`

	err := r.db.QueryRow(query, postSubmissionUUID).Scan(&total_ratings, &sum_rating)
	if err != nil {
		return nil, err
	}

	if total_ratings == 0 {
		return &total_ratings, nil
	}

	rating := int(math.RoundToEven(float64(sum_rating) / float64(total_ratings)))
	return &rating, nil
}

func (r *PostRepos) RateSubmissionPost(raterId uint64, submissionUUID string, rating uint64) error {
	var count int

	err := r.db.QueryRow("SELECT COUNT(*) FROM posts_submissions_ratings WHERE post_submission_uuid = ? AND rater_id = ?", submissionUUID, raterId).Scan(&count)
	if err != nil {
		return err
	}

	if count != 0 {
		_, err := r.db.
			Query("UPDATE posts_submissions_ratings SET rating = ? WHERE post_submission_uuid = ? AND rater_id = ?",
				rating, submissionUUID, raterId)

		return err
	}

	_, err = r.db.
		Query("INSERT INTO posts_submissions_ratings (post_submission_uuid, rater_id, rating) VALUES (?, ?, ?)", submissionUUID, raterId, rating)

	return err
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

	submissions, err := r.GetSubmissions(uuid)
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
		submissions, _ := r.GetSubmissions(post.UUID)

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

func (r *PostRepos) GetComment(uuid string) (*models.Comment, error) {
	var comment models.Comment

	query := `
		SELECT id, uuid, post_uuid, author_id, comment, deleted, created 
		FROM posts_comments WHERE deleted = 0 AND uuid = ?
  	`

	err := r.db.
		QueryRow(query, uuid).
		Scan(
			&comment.ID,
			&comment.UUID,
			&comment.PostUUID,
			&comment.AuthorID,
			&comment.Comment,
			&comment.Deleted,
			&comment.Created,
		)

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *PostRepos) GetIdeaComments(postUUID string) (models.Comments, error) {
	var comments = models.Comments{}

	query := `
		SELECT id, uuid, post_uuid, author_id, comment, deleted, created 
		FROM posts_comments WHERE deleted = 0 AND post_uuid = ? 
		ORDER BY id DESC
  	`

	results, err := r.db.Query(query, postUUID)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var comment models.Comment

		results.Scan(
			&comment.ID,
			&comment.UUID,
			&comment.PostUUID,
			&comment.AuthorID,
			&comment.Comment,
			&comment.Deleted,
			&comment.Created,
		)

		comments = append(comments, comment)
	}
	return comments, nil
}

func (r *PostRepos) CreateIdeaComment(authorId uint64, postUUID string, comment string) (string, error) {
	uuid := uuid.NewV4().String()
	query := `
		INSERT INTO posts_comments (uuid, post_uuid, author_id, comment) 
		VALUES (?, ?, ?, ?)
	`
	_, err := r.db.Query(query, uuid, postUUID, authorId, comment)

	return uuid, err
}

func (r *PostRepos) GetIdeaSubmission(submissionUUID string) (*models.IdeaSubmission, error) {
	submission, err := r.GetSubmission(submissionUUID)
	if err != nil {
		return nil, err
	}

	rating, err := r.GetSubmissionRating(submissionUUID)
	if err != nil {
		return nil, err
	}

	return &models.IdeaSubmission{
		ID:       submission.ID,
		UUID:     submission.UUID,
		PostUUID: submission.PostUUID,
		AuthorID: submission.AuthorID,
		Comment:  submission.Comment,
		Rating:   *rating,
		Deleted:  submission.Deleted,
		Modified: submission.Modified,
		Created:  submission.Created,
	}, nil
}

func (r *PostRepos) GetIdeaSubmissions(postUUID string) (models.IdeaSubmissions, error) {
	var idea_submissions = models.IdeaSubmissions{}

	submissions, err := r.GetSubmissions(postUUID)
	if err != nil {
		return nil, err
	}

	if len(submissions) == 0 {
		return idea_submissions, nil
	}

	// TODO make more performant
	for _, submission := range submissions {
		rating, _ := r.GetSubmissionRating(submission.UUID)

		idea_submissions = append(idea_submissions,
			models.IdeaSubmission{
				ID:       submission.ID,
				UUID:     submission.UUID,
				PostUUID: submission.PostUUID,
				AuthorID: submission.AuthorID,
				Comment:  submission.Comment,
				Rating:   *rating,
				Deleted:  submission.Deleted,
				Modified: submission.Modified,
				Created:  submission.Created,
			})
	}

	return idea_submissions, nil
}
