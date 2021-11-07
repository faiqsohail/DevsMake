package handlers

import (
	"devsmake/models"
	"devsmake/persistence/interfaces"
	"devsmake/restapi/operations/idea_submissions"
	"devsmake/util"
	"strings"

	"github.com/go-openapi/runtime/middleware"
)

type IdeaPutSubmissionRatingHandler struct {
	accountRepo interfaces.AccountRepository
	postRepos   interfaces.PostRepositories
}

func NewIdeaPutSubmissionRatingHandler(accountRepo interfaces.AccountRepository, postRepos interfaces.PostRepositories) *IdeaPutSubmissionRatingHandler {
	return &IdeaPutSubmissionRatingHandler{
		accountRepo: accountRepo,
		postRepos:   postRepos,
	}
}

func (handler *IdeaPutSubmissionRatingHandler) Handle(params idea_submissions.PutIdeasUUIDSubmissionsRateParams, principal *models.Principal) middleware.Responder {
	user, err := util.FetchAuthedUser(string(*principal))

	if err != nil {
		errMsg := err.Error()
		return idea_submissions.NewPutIdeasUUIDSubmissionsRateDefault(500).WithPayload(
			&models.Error{
				Message: &errMsg,
			},
		)
	}

	_, err = handler.postRepos.GetPost(params.UUID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			errMsg := "no idea post with the specified uuid exists"
			return idea_submissions.NewPutIdeasUUIDSubmissionsRateDefault(404).WithPayload(
				&models.Error{
					Message: &errMsg,
				},
			)
		} else {
			errMsg := err.Error()
			return idea_submissions.NewPutIdeasUUIDSubmissionsRateDefault(500).WithPayload(
				&models.Error{
					Message: &errMsg,
				},
			)
		}
	}

	_, err = handler.postRepos.GetSubmission(*params.Rating.SubmissionUUID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			errMsg := "no submission post with the specified uuid exists"
			return idea_submissions.NewPutIdeasUUIDSubmissionsRateDefault(404).WithPayload(
				&models.Error{
					Message: &errMsg,
				},
			)
		} else {
			errMsg := err.Error()
			return idea_submissions.NewPutIdeasUUIDSubmissionsRateDefault(500).WithPayload(
				&models.Error{
					Message: &errMsg,
				},
			)
		}
	}

	storedUser, _ := handler.accountRepo.GetUser(uint64(*user.ID), true)

	err = handler.postRepos.RateSubmissionPost(storedUser.ID, *params.Rating.SubmissionUUID, uint64(*params.Rating.Rating))

	if err != nil {
		errMsg := err.Error()
		return idea_submissions.NewPutIdeasUUIDSubmissionsRateDefault(500).WithPayload(
			&models.Error{
				Message: &errMsg,
			},
		)
	}

	submission, _ := handler.postRepos.GetIdeaSubmission(*params.Rating.SubmissionUUID)
	return idea_submissions.NewPutIdeasUUIDSubmissionsRateOK().WithPayload(submission.PublicIdeaSubmission())
}
