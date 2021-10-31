package handlers

import (
	"devsmake/models"
	"devsmake/persistence/interfaces"
	"devsmake/restapi/operations/profile"
	"devsmake/util"

	"github.com/go-openapi/runtime/middleware"
)

type ProfileIDHandler struct {
	db interfaces.AccountRepository
}

func NewProfileIDHandler(accountRepo interfaces.AccountRepository) *ProfileIDHandler {
	return &ProfileIDHandler{
		db: accountRepo,
	}
}

func (handler *ProfileIDHandler) Handle(params profile.GetProfileIDParams) middleware.Responder {
	storedUser, err := handler.db.GetUser(uint64(params.ID))
	if err != nil {
		msg := err.Error()
		return profile.NewGetProfileIDDefault(500).WithPayload(
			&models.Error{
				Message: &msg,
			},
		)
	}

	return profile.NewGetProfileOK().WithPayload(
		&models.Profile{
			Identifier: int64(storedUser.ID),
			Username:   storedUser.Username,
			Points:     int64(storedUser.Points),
			AvatarURL:  util.GenerateAvatarUrl(storedUser.Username),
		},
	)
}
