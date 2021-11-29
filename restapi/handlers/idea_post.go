package handlers

import (
	"devsmake/models"
	"devsmake/persistence/interfaces"
	"devsmake/restapi/operations/idea_post"
	"devsmake/util"

	"github.com/go-openapi/runtime/middleware"
)

type IdeaCreateHandler struct {
	accountRepo interfaces.AccountRepository
	postRepos   interfaces.PostRepositories
}

func NewIdeaCreateHandler(accountRepo interfaces.AccountRepository, postRepos interfaces.PostRepositories) *IdeaCreateHandler {
	return &IdeaCreateHandler{
		accountRepo: accountRepo,
		postRepos:   postRepos,
	}
}

func (handler *IdeaCreateHandler) Handle(params idea_post.PostIdeasParams, principal *models.Principal) middleware.Responder {
	user, err := util.FetchAuthedUser(string(*principal))

	if err != nil {
		errMsg := err.Error()
		return idea_post.NewPostIdeasDefault(500).WithPayload(
			&models.Error{
				Message: &errMsg,
			},
		)
	}

	storedUser, _ := handler.accountRepo.GetUser(uint64(*user.ID), true)

	uuid, err := handler.postRepos.CreatePost(storedUser.ID, *params.Idea.Title, *params.Idea.Description)
	if err != nil {
		errMsg := err.Error()
		return idea_post.NewPostIdeasDefault(500).WithPayload(
			&models.Error{
				Message: &errMsg,
			},
		)
	}

	createdIdea, _ := handler.postRepos.GetIdea(uuid)
	handler.accountRepo.AddPoints(int(*user.ID), 2)

	return idea_post.NewPostIdeasOK().WithPayload(createdIdea.PublicIdea())
}
