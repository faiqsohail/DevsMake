package handlers

import (
	"devsmake/models"
	"devsmake/persistence/interfaces"
	"devsmake/restapi/operations/profile"

	"github.com/go-openapi/runtime/middleware"
)

type ProfilesHandler struct {
	db interfaces.AccountRepository
}

func NewProfilesHandler(accountRepo interfaces.AccountRepository) *ProfilesHandler {
	return &ProfilesHandler{
		db: accountRepo,
	}
}

func (handler *ProfilesHandler) Handle(params profile.GetProfilesParams) middleware.Responder {
	users, err := handler.db.GetUsers(uint64(params.Limit), uint64(params.Offset), *params.Sort)
	if err != nil {
		errMsg := err.Error()
		return profile.NewGetProfilesDefault(500).WithPayload(
			&models.Error{
				Message: &errMsg,
			},
		)
	}
	return profile.NewGetProfilesOK().WithPayload(users.PublicUsers())
}
