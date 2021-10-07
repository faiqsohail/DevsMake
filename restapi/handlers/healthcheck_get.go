package handlers

import (
	"devsmake/persistence"
	"devsmake/restapi/operations/general"

	"github.com/go-openapi/runtime/middleware"
)

type HealthCheckHandler struct {
	db *persistence.Repositories
}

func NewHealthCheckHandler(db *persistence.Repositories) *HealthCheckHandler {
	return &HealthCheckHandler{
		db: db,
	}
}

func (handler *HealthCheckHandler) Handle(params general.GetHealthcheckParams) middleware.Responder {
	err := handler.db.Ping()

	if err != nil {
		return general.NewGetHealthcheckInternalServerError().WithPayload(&general.GetHealthcheckInternalServerErrorBody{
			Status: "unavailable",
		})
	}
	return general.NewGetHealthcheckOK().WithPayload(&general.GetHealthcheckOKBody{
		Status: "available",
	})
}
