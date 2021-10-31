package handlers

import (
	"context"
	"devsmake/models"
	"devsmake/persistence/interfaces"
	"devsmake/restapi/operations/profile"
	"devsmake/util"

	"github.com/go-openapi/runtime/middleware"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
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
	p := string(*principal)
	oauthClient := util.GetOAuthConfig().Client(context.TODO(), &oauth2.Token{AccessToken: p})
	client := github.NewClient(oauthClient)

	user, _, err := client.Users.Get(context.TODO(), "")
	if err != nil {
		errMsg := "Unable to fetch the logged in user"
		return profile.NewGetProfileDefault(500).WithPayload(
			&models.Error{
				Message: &errMsg,
			},
		)
	}

	storedUser, err := handler.db.GetUser(uint64(*user.ID))
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
