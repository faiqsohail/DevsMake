package handlers

import (
	"devsmake/models"
	"devsmake/persistence/interfaces"
	"devsmake/restapi/operations/idea_submissions"
	"strings"

	"github.com/go-openapi/runtime/middleware"
)

type IdeaGetSubmissionsHandler struct {
	db interfaces.PostRepositories
}

func NewIdeaGetSubmissionsHandler(postRepos interfaces.PostRepositories) *IdeaGetSubmissionsHandler {
	return &IdeaGetSubmissionsHandler{
		db: postRepos,
	}
}

func (handler *IdeaGetSubmissionsHandler) Handle(params idea_submissions.GetIdeasUUIDSubmissionsParams, principal *models.Principal) middleware.Responder {
	_, err := handler.db.GetPost(params.UUID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			errMsg := "no idea post with the specified uuid exists"
			return idea_submissions.NewGetIdeasUUIDSubmissionsDefault(404).WithPayload(
				&models.Error{
					Message: &errMsg,
				},
			)
		} else {
			errMsg := err.Error()
			return idea_submissions.NewGetIdeasUUIDSubmissionsDefault(500).WithPayload(
				&models.Error{
					Message: &errMsg,
				},
			)
		}
	}

	submissions, err := handler.db.GetIdeaSubmissions(params.UUID)
	if err != nil {
		errMsg := err.Error()
		return idea_submissions.NewGetIdeasUUIDSubmissionsDefault(500).WithPayload(
			&models.Error{
				Message: &errMsg,
			},
		)
	}

	return idea_submissions.NewGetIdeasUUIDSubmissionsOK().WithPayload(submissions.PublicIdeaSubmissions())
}
