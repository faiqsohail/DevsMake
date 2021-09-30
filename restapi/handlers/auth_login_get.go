package handlers

import (
	"devsmake/restapi/operations/auth"
	"devsmake/util"

	"github.com/go-openapi/runtime/middleware"
)

type AuthLoginHandler struct{}

func NewAuthLoginHandler() *AuthLoginHandler {
	return &AuthLoginHandler{}
}

func (handler *AuthLoginHandler) Handle(params auth.GetAuthLoginParams) middleware.Responder {
	return util.DoLogin(params.HTTPRequest)
}
