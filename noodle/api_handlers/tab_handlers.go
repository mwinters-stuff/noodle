package api_handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
)

var (
	RegisterTabApiHandlers = RegisterTabApiHandlersImpl
)

func HandlerTabGet(db database.Database, params noodle_api.GetNoodleTabsParams, principal *models.Principal) middleware.Responder {
	tabs, err := db.Tables().TabTable().GetAll()
	if err != nil {
		return noodle_api.NewGetNoodleTabsConflict().WithPayload(&models.Error{Message: err.Error()})
	}
	return noodle_api.NewGetNoodleTabsOK().WithPayload(tabs)
}

func HandlerTabPost(db database.Database, params noodle_api.PostNoodleTabsParams, principal *models.Principal) middleware.Responder {
	if params.Action == "insert" {
		err := db.Tables().TabTable().Insert(params.Tab)
		if err != nil {
			return noodle_api.NewPostNoodleTabsConflict().WithPayload(&models.Error{Message: err.Error()})
		}
	}
	if params.Action == "update" {
		err := db.Tables().TabTable().Update(*params.Tab)
		if err != nil {
			return noodle_api.NewPostNoodleTabsConflict().WithPayload(&models.Error{Message: err.Error()})
		}
	}

	return noodle_api.NewPostNoodleTabsOK().WithPayload(params.Tab)
}

func HandlerTabDelete(db database.Database, params noodle_api.DeleteNoodleTabsParams, principal *models.Principal) middleware.Responder {
	err := db.Tables().TabTable().Delete(params.Tabid)
	if err != nil {
		return noodle_api.NewDeleteNoodleTabsConflict().WithPayload(&models.Error{Message: err.Error()})
	}
	return noodle_api.NewDeleteNoodleTabsOK()

}

func RegisterTabApiHandlersImpl(api *operations.NoodleAPI, db database.Database) {
	api.NoodleAPIGetNoodleTabsHandler = noodle_api.GetNoodleTabsHandlerFunc(func(params noodle_api.GetNoodleTabsParams, principal *models.Principal) middleware.Responder {
		return HandlerTabGet(db, params, principal)
	})
	api.NoodleAPIPostNoodleTabsHandler = noodle_api.PostNoodleTabsHandlerFunc(func(params noodle_api.PostNoodleTabsParams, principal *models.Principal) middleware.Responder {
		return HandlerTabPost(db, params, principal)
	})
	api.NoodleAPIDeleteNoodleTabsHandler = noodle_api.DeleteNoodleTabsHandlerFunc(func(params noodle_api.DeleteNoodleTabsParams, principal *models.Principal) middleware.Responder {
		return HandlerTabDelete(db, params, principal)
	})

}
