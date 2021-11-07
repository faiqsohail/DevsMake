package interfaces

import "devsmake/persistence/models"

type PostRating int64

const (
	NoRating PostRating = 0
	Dislike  PostRating = -1
	Like     PostRating = 1
)

type PostRepositories interface {
	GetPost(string) (*models.Post, error)
	GetPosts(uint64, uint64, string) (models.Posts, error)
	CreatePost(uint64, string, string) (string, error)
	GetIdea(string) (*models.Idea, error)
	GetIdeas(uint64, uint64, string) (models.Ideas, error)
	GetPostRatings(string, PostRating) (*int, error)
	RatePost(uint64, string, PostRating) error
	GetComment(string) (*models.Comment, error)
	GetIdeaComments(string) (models.Comments, error)
	CreateIdeaComment(uint64, string, string) (string, error)
	CreateSubmission(uint64, string, string) (string, error)
	GetSubmission(string) (*models.Submission, error)
	GetSubmissions(string) (models.Submissions, error)
	GetSubmissionRating(string) (*int, error)
	GetIdeaSubmission(string) (*models.IdeaSubmission, error)
	GetIdeaSubmissions(string) (models.IdeaSubmissions, error)
}
