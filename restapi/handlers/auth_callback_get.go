package handlers

import (
	"devsmake/restapi/operations/auth"
	"devsmake/util"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

type AuthCallbackHandler struct{}

func NewAuthCallbackHandler() *AuthCallbackHandler {
	return &AuthCallbackHandler{}
}

func (handler *AuthCallbackHandler) Handle(params auth.GetAuthCallbackParams) middleware.Responder {
	token, err := util.DoCallback(params.HTTPRequest)
	if err != nil {
		return middleware.NotImplemented("Unexpected error occured")
	}

	return middleware.ResponderFunc(
		func(w http.ResponseWriter, pr runtime.Producer) {
			sessionCookie := http.Cookie{Name: "sessionCookie", Value: token.AccessToken, Expires: token.Expiry}
			http.SetCookie(w, &sessionCookie)

			http.Redirect(w, params.HTTPRequest, "/", http.StatusTemporaryRedirect)
		})
}
