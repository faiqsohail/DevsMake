package handlers

import (
	"devsmake/models"
	"devsmake/persistence/interfaces"
	"devsmake/restapi/operations/idea_post"

	"github.com/go-openapi/runtime/middleware"
)

type IdeasHandler struct {
	db interfaces.PostRepositories
}

func NewIdeasHandler(postRepos interfaces.PostRepositories) *IdeasHandler {
	return &IdeasHandler{
		db: postRepos,
	}
}

func (handler *IdeasHandler) Handle(params idea_post.GetIdeasParams) middleware.Responder {
	query := ""
	if params.Query != nil {
		query = *params.Query
	}
	ideas, err := handler.db.GetIdeas(uint64(*params.Limit), uint64(*params.Offset), query)
	if err != nil {
		errMsg := err.Error()
		return idea_post.NewGetIdeasDefault(500).WithPayload(
			&models.Error{
				Message: &errMsg,
			},
		)
	}
	return idea_post.NewGetIdeasOK().WithPayload(ideas.PublicIdeas())
}
