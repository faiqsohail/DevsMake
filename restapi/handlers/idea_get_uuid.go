package handlers

import (
	"devsmake/models"
	"devsmake/persistence/interfaces"
	"devsmake/restapi/operations/idea_post"

	"github.com/go-openapi/runtime/middleware"
)

type IdeasUUIDHandler struct {
	db interfaces.PostRepositories
}

func NewIdeasUUIDHandler(postRepos interfaces.PostRepositories) *IdeasUUIDHandler {
	return &IdeasUUIDHandler{
		db: postRepos,
	}
}

func (handler *IdeasUUIDHandler) Handle(params idea_post.GetIdeasUUIDParams) middleware.Responder {
	idea, err := handler.db.GetIdea(params.UUID)
	if err != nil {
		msg := err.Error()
		return idea_post.NewGetIdeasUUIDDefault(500).WithPayload(
			&models.Error{
				Message: &msg,
			},
		)
	}
	return idea_post.NewGetIdeasUUIDOK().WithPayload(idea.PublicIdea())
}
