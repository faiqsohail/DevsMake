package handlers

import (
	"devsmake/models"
	"devsmake/persistence/interfaces"
	"devsmake/restapi/operations/idea_comments"
	"devsmake/util"
	"strings"

	"github.com/go-openapi/runtime/middleware"
)

type IdeaCreateCommentsHandler struct {
	accountRepo interfaces.AccountRepository
	postRepos   interfaces.PostRepositories
}

func NewIdeaCreateCommentsHandler(accountRepo interfaces.AccountRepository, postRepos interfaces.PostRepositories) *IdeaCreateCommentsHandler {
	return &IdeaCreateCommentsHandler{
		accountRepo: accountRepo,
		postRepos:   postRepos,
	}
}

func (handler *IdeaCreateCommentsHandler) Handle(params idea_comments.PostIdeasUUIDCommentsParams, principal *models.Principal) middleware.Responder {
	user, err := util.FetchAuthedUser(string(*principal))

	if err != nil {
		errMsg := err.Error()
		return idea_comments.NewPostIdeasUUIDCommentsDefault(500).WithPayload(
			&models.Error{
				Message: &errMsg,
			},
		)
	}

	_, err = handler.postRepos.GetPost(params.UUID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			errMsg := "no idea post with the specified uuid exists"
			return idea_comments.NewPostIdeasUUIDCommentsDefault(404).WithPayload(
				&models.Error{
					Message: &errMsg,
				},
			)
		} else {
			errMsg := err.Error()
			return idea_comments.NewPostIdeasUUIDCommentsDefault(500).WithPayload(
				&models.Error{
					Message: &errMsg,
				},
			)
		}
	}

	storedUser, _ := handler.accountRepo.GetUser(uint64(*user.ID), true)

	uuid, err := handler.postRepos.CreateIdeaComment(storedUser.ID, params.UUID, *params.Comment.Comment)
	if err != nil {
		errMsg := err.Error()
		return idea_comments.NewPostIdeasUUIDCommentsDefault(500).WithPayload(
			&models.Error{
				Message: &errMsg,
			},
		)
	}

	comment, _ := handler.postRepos.GetComment(uuid)
	return idea_comments.NewPostIdeasUUIDCommentsOK().WithPayload(comment.PublicComment())
}
