package handlers

import (
	"devsmake/models"
	"devsmake/persistence/interfaces"
	"devsmake/restapi/operations/profile"
	"devsmake/util"

	"github.com/go-openapi/runtime/middleware"
)

type ProfileHandler struct {
	db interfaces.AccountRepository
}

func NewProfileHandler(accountRepo interfaces.AccountRepository) *ProfileHandler {
	return &ProfileHandler{
		db: accountRepo,
	}
}

func (handler *ProfileHandler) Handle(params profile.GetProfileParams, principal *models.Principal) middleware.Responder {
	user, err := util.FetchAuthedUser(string(*principal))

	if err != nil {
		errMsg := err.Error()
		return profile.NewGetProfileDefault(500).WithPayload(
			&models.Error{
				Message: &errMsg,
			},
		)
	}

	storedUser, err := handler.db.GetUser(uint64(*user.ID), true)
	if err != nil {
		errMsg := err.Error()
		return profile.NewGetProfileDefault(500).WithPayload(
			&models.Error{
				Message: &errMsg,
			},
		)
	}

	return profile.NewGetProfileOK().WithPayload(storedUser.PublicUser())
}
