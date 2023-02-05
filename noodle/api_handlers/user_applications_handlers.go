package api_handlers

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
)

var (
	RegisterUserApplicationsApiHandlers = RegisterUserApplicationsApiHandlersImpl
)

func HandlerUserApplicationGet(db database.Database, params noodle_api.GetNoodleUserApplicationsParams, principal *models.Principal) middleware.Responder {
	applicationTabs, err := db.Tables().UserApplicationsTable().GetUserApps(params.UserID)
	if err != nil {
		return noodle_api.NewGetNoodleUserApplicationsConflict().WithPayload(&models.Error{Message: err.Error()})
	}
	return noodle_api.NewGetNoodleUserApplicationsOK().WithPayload(applicationTabs)
}

func HandlerUserAllowedApplicationGet(db database.Database, params noodle_api.GetNoodleUserAllowedApplicationsParams, principal *models.Principal) middleware.Responder {
	applications, err := db.Tables().UserApplicationsTable().GetUserAllowdApplications(params.UserID)
	if err != nil {
		return noodle_api.NewGetNoodleUserAllowedApplicationsConflict().WithPayload(&models.Error{Message: err.Error()})
	}
	return noodle_api.NewGetNoodleUserAllowedApplicationsOK().WithPayload(applications)
}

func HandlerUserApplicationPost(db database.Database, params noodle_api.PostNoodleUserApplicationsParams, principal *models.Principal) middleware.Responder {
	err := db.Tables().UserApplicationsTable().Insert(params.UserApplication)

	if err != nil {
		return noodle_api.NewPostNoodleTabsConflict().WithPayload(&models.Error{Message: err.Error()})
	}

	return noodle_api.NewPostNoodleUserApplicationsOK().WithPayload(params.UserApplication)
}

func HandlerUserApplicationDelete(db database.Database, params noodle_api.DeleteNoodleUserApplicationsParams, principal *models.Principal) middleware.Responder {
	err := db.Tables().UserApplicationsTable().Delete(params.UserApplicationID)
	if err != nil {
		return noodle_api.NewDeleteNoodleUserApplicationsConflict().WithPayload(&models.Error{Message: err.Error()})
	}
	return noodle_api.NewDeleteNoodleUserApplicationsOK()
}

func RegisterUserApplicationsApiHandlersImpl(api *operations.NoodleAPI, db database.Database) {
	api.NoodleAPIGetNoodleUserApplicationsHandler = noodle_api.GetNoodleUserApplicationsHandlerFunc(func(params noodle_api.GetNoodleUserApplicationsParams, principal *models.Principal) middleware.Responder {
		return HandlerUserApplicationGet(db, params, principal)
	})
	api.NoodleAPIPostNoodleUserApplicationsHandler = noodle_api.PostNoodleUserApplicationsHandlerFunc(func(params noodle_api.PostNoodleUserApplicationsParams, principal *models.Principal) middleware.Responder {
		return HandlerUserApplicationPost(db, params, principal)
	})
	api.NoodleAPIDeleteNoodleUserApplicationsHandler = noodle_api.DeleteNoodleUserApplicationsHandlerFunc(func(params noodle_api.DeleteNoodleUserApplicationsParams, principal *models.Principal) middleware.Responder {
		return HandlerUserApplicationDelete(db, params, principal)
	})
	api.NoodleAPIGetNoodleUserAllowedApplicationsHandler = noodle_api.GetNoodleUserAllowedApplicationsHandlerFunc(func(params noodle_api.GetNoodleUserAllowedApplicationsParams, principal *models.Principal) middleware.Responder {
		return HandlerUserAllowedApplicationGet(db, params, principal)
	})
}
