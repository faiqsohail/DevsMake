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
	GetIdea(string) (*models.Idea, error)
	GetIdeas(uint64, uint64, string) (models.Ideas, error)
	GetPostSubmissions(string) (models.Submissions, error)
	GetPostRatings(string, PostRating) (*int, error)
}