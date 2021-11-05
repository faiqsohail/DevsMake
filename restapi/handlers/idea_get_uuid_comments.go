package handlers

import (
	"devsmake/models"
	"devsmake/persistence/interfaces"
	"devsmake/restapi/operations/idea_comments"
	"strings"

	"github.com/go-openapi/runtime/middleware"
)

type IdeaGetCommentsHandler struct {
	db interfaces.PostRepositories
}

func NewIdeaGetCommentsHandler(postRepos interfaces.PostRepositories) *IdeaGetCommentsHandler {
	return &IdeaGetCommentsHandler{
		db: postRepos,
	}
}

func (handler *IdeaGetCommentsHandler) Handle(params idea_comments.GetIdeasUUIDCommentsParams) middleware.Responder {
	_, err := handler.db.GetPost(params.UUID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			errMsg := "no idea post with the specified uuid exists"
			return idea_comments.NewGetIdeasUUIDCommentsDefault(404).WithPayload(
				&models.Error{
					Message: &errMsg,
				},
			)
		} else {
			errMsg := err.Error()
			return idea_comments.NewGetIdeasUUIDCommentsDefault(500).WithPayload(
				&models.Error{
					Message: &errMsg,
				},
			)
		}
	}

	comments, err := handler.db.GetIdeaComments(params.UUID)
	if err != nil {
		errMsg := err.Error()
		return idea_comments.NewGetIdeasUUIDCommentsDefault(500).WithPayload(
			&models.Error{
				Message: &errMsg,
			},
		)
	}
	return idea_comments.NewGetIdeasUUIDCommentsOK().WithPayload(comments.PublicComments())
}
