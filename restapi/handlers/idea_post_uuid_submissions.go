package handlers

import (
	"devsmake/models"
	"devsmake/persistence/interfaces"
	"devsmake/restapi/operations/idea_submissions"
	"devsmake/util"
	"strings"

	"github.com/go-openapi/runtime/middleware"
)

type IdeaCreateSubmissionsHandler struct {
	accountRepo interfaces.AccountRepository
	postRepos   interfaces.PostRepositories
}

func NewIdeaCreateSubmissionsHandler(accountRepo interfaces.AccountRepository, postRepos interfaces.PostRepositories) *IdeaCreateSubmissionsHandler {
	return &IdeaCreateSubmissionsHandler{
		accountRepo: accountRepo,
		postRepos:   postRepos,
	}
}

func (handler *IdeaCreateSubmissionsHandler) Handle(params idea_submissions.PostIdeasUUIDSubmissionsParams, principal *models.Principal) middleware.Responder {
	user, err := util.FetchAuthedUser(string(*principal))

	if err != nil {
		errMsg := err.Error()
		return idea_submissions.NewPostIdeasUUIDSubmissionsDefault(500).WithPayload(
			&models.Error{
				Message: &errMsg,
			},
		)
	}

	_, err = handler.postRepos.GetPost(params.UUID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			errMsg := "no idea post with the specified uuid exists"
			return idea_submissions.NewPostIdeasUUIDSubmissionsDefault(404).WithPayload(
				&models.Error{
					Message: &errMsg,
				},
			)
		} else {
			errMsg := err.Error()
			return idea_submissions.NewPostIdeasUUIDSubmissionsDefault(500).WithPayload(
				&models.Error{
					Message: &errMsg,
				},
			)
		}
	}

	storedUser, _ := handler.accountRepo.GetUser(uint64(*user.ID), true)

	uuid, err := handler.postRepos.CreateSubmission(storedUser.ID, params.UUID, *params.Submision.Comment)
	if err != nil {
		errMsg := err.Error()
		return idea_submissions.NewPostIdeasUUIDSubmissionsDefault(500).WithPayload(
			&models.Error{
				Message: &errMsg,
			},
		)
	}

	handler.accountRepo.AddPoints(int(*user.ID), 10)

	submission, _ := handler.postRepos.GetIdeaSubmission(uuid)
	return idea_submissions.NewPostIdeasUUIDSubmissionsOK().WithPayload(submission.PublicIdeaSubmission())
}
