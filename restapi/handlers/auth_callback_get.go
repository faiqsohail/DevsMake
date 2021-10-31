package handlers

import (
	"context"
	"devsmake/persistence/interfaces"
	"devsmake/restapi/operations/auth"
	"devsmake/util"
	"net/http"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/go-github/github"
)

type AuthCallbackHandler struct {
	db interfaces.AccountRepository
}

func NewAuthCallbackHandler(accountRepo interfaces.AccountRepository) *AuthCallbackHandler {
	return &AuthCallbackHandler{
		db: accountRepo,
	}
}

func (handler *AuthCallbackHandler) Handle(params auth.GetAuthCallbackParams) middleware.Responder {
	token, err := util.DoCallback(params.HTTPRequest)
	if err != nil {
		return middleware.NotImplemented("Unexpected error occured")
	}

	oauthClient := util.GetOAuthConfig().Client(context.TODO(), token)
	client := github.NewClient(oauthClient)

	user, _, err := client.Users.Get(context.TODO(), "")
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, "Unable to fetch the logged in user")
	}

	// Create the user account if it doesn't exist
	_, err = handler.db.GetUser(uint64(*user.ID), true)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			err := handler.db.CreateUser(uint64(*user.ID), *user.Login)
			if err != nil {
				return middleware.Error(http.StatusInternalServerError, "Unable to save user information")
			}
		}
		return middleware.Error(http.StatusInternalServerError, "Unable to determine if user exists")
	}

	return middleware.ResponderFunc(
		func(w http.ResponseWriter, pr runtime.Producer) {
			sessionCookie := http.Cookie{
				Name:    "sessionCookie",
				Value:   token.AccessToken,
				Expires: token.Expiry,
				Path:    "/",
				Secure:  true,
			}
			http.SetCookie(w, &sessionCookie)

			http.Redirect(w, params.HTTPRequest, "/", http.StatusTemporaryRedirect)
		})
}
