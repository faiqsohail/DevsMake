package handlers

import (
	"devsmake/models"
	"devsmake/persistence/interfaces"
	"devsmake/restapi/operations/idea_post"
	"devsmake/util"
	"strings"

	"github.com/go-openapi/runtime/middleware"
)

type IdeaPutRatingHandler struct {
	accountRepo interfaces.AccountRepository
	postRepos   interfaces.PostRepositories
}

func NewIdeaPutRatingHandler(accountRepo interfaces.AccountRepository, postRepos interfaces.PostRepositories) *IdeaPutRatingHandler {
	return &IdeaPutRatingHandler{
		accountRepo: accountRepo,
		postRepos:   postRepos,
	}
}

func (handler *IdeaPutRatingHandler) Handle(params idea_post.PutIdeasUUIDRateParams, principal *models.Principal) middleware.Responder {
	user, err := util.FetchAuthedUser(string(*principal))

	if err != nil {
		errMsg := err.Error()
		return idea_post.NewPutIdeasUUIDRateDefault(500).WithPayload(
			&models.Error{
				Message: &errMsg,
			},
		)
	}

	_, err = handler.postRepos.GetPost(params.UUID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			errMsg := "no idea post with the specified uuid exists"
			return idea_post.NewPutIdeasUUIDRateDefault(404).WithPayload(
				&models.Error{
					Message: &errMsg,
				},
			)
		} else {
			errMsg := err.Error()
			return idea_post.NewPutIdeasUUIDRateDefault(500).WithPayload(
				&models.Error{
					Message: &errMsg,
				},
			)
		}
	}

	storedUser, _ := handler.accountRepo.GetUser(uint64(*user.ID), true)

	var postRating interfaces.PostRating

	switch rating := *params.Rating.Rating; rating {
	case idea_post.PutIdeasUUIDRateBodyRatingLike:
		postRating = interfaces.Like
	case idea_post.PutIdeasUUIDRateBodyRatingDislike:
		postRating = interfaces.Dislike
	default:
		postRating = interfaces.NoRating
	}

	err = handler.postRepos.RatePost(storedUser.ID, params.UUID, postRating)
	if err != nil {
		errMsg := err.Error()
		return idea_post.NewPutIdeasUUIDRateDefault(500).WithPayload(
			&models.Error{
				Message: &errMsg,
			},
		)
	}
	idea, _ := handler.postRepos.GetIdea(params.UUID)
	return idea_post.NewPutIdeasUUIDRateOK().WithPayload(idea.PublicIdea())
}
