package handlers

import (
	"context"
	"devsmake/models"
	"devsmake/restapi/operations/profile"
	"devsmake/util"
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type ProfileHandler struct{}

func NewProfileHandler() *ProfileHandler {
	return &ProfileHandler{}
}

func (handler *ProfileHandler) Handle(params profile.GetProfileParams, principal *models.Principal) middleware.Responder {
	p := string(*principal)
	oauthClient := util.GetOAuthConfig().Client(context.TODO(), &oauth2.Token{AccessToken: p})
	client := github.NewClient(oauthClient)

	user, _, err := client.Users.Get(context.TODO(), "")
	if err != nil {
		fmt.Printf("client.Users.Get() faled with '%s'\n", err)
		return middleware.NotImplemented("Unable to fetch the logged in user")
	}

	return profile.NewGetProfileOK().WithPayload(
		&models.Profile{
			Identifier: *user.ID,
			Username:   *user.Login,
		},
	)
}
