package api_handlers

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
)

var (
	RegisterApplicationsApiHandlers = RegisterApplicationsApiHandlersImpl
)

func HandlerApplicationGet(db database.Database, params noodle_api.GetNoodleApplicationsParams, principal *models.Principal) middleware.Responder {
	if params.ApplicationID != nil && params.ApplicationTemplate != nil {
		return noodle_api.NewGetNoodleApplicationsConflict().WithPayload(&models.Error{Message: "incorrect parameters - both supplied"})
	}
	if params.ApplicationID != nil {
		application, err := db.Tables().ApplicationsTable().GetID(*params.ApplicationID)
		if err != nil {
			return noodle_api.NewGetNoodleApplicationsConflict().WithPayload(&models.Error{Message: err.Error()})
		}
		return noodle_api.NewGetNoodleApplicationsOK().WithPayload([]*models.Application{&application})
	}
	if params.ApplicationTemplate != nil {
		applications, err := db.Tables().ApplicationsTable().GetTemplateID(*params.ApplicationTemplate)
		if err != nil {
			return noodle_api.NewGetNoodleApplicationsConflict().WithPayload(&models.Error{Message: err.Error()})
		}
		return noodle_api.NewGetNoodleApplicationsOK().WithPayload(applications)
	}
	return noodle_api.NewGetNoodleApplicationsConflict().WithPayload(&models.Error{Message: "incorrect parameters - both nil"})
}

func HandlerApplicationPost(db database.Database, params noodle_api.PostNoodleApplicationsParams, principal *models.Principal) middleware.Responder {
	err := db.Tables().ApplicationsTable().Insert(params.Application)

	if err != nil {
		return noodle_api.NewPostNoodleTabsConflict().WithPayload(&models.Error{Message: err.Error()})
	}

	return noodle_api.NewPostNoodleApplicationsOK().WithPayload(params.Application)
}

func HandlerApplicationDelete(db database.Database, params noodle_api.DeleteNoodleApplicationsParams, principal *models.Principal) middleware.Responder {
	err := db.Tables().ApplicationsTable().Delete(params.ApplicationID)
	if err != nil {
		return noodle_api.NewDeleteNoodleApplicationsConflict().WithPayload(&models.Error{Message: err.Error()})
	}
	return noodle_api.NewDeleteNoodleApplicationsOK()

}

func RegisterApplicationsApiHandlersImpl(api *operations.NoodleAPI, db database.Database) {
	api.NoodleAPIGetNoodleApplicationsHandler = noodle_api.GetNoodleApplicationsHandlerFunc(func(params noodle_api.GetNoodleApplicationsParams, principal *models.Principal) middleware.Responder {
		return HandlerApplicationGet(db, params, principal)
	})
	api.NoodleAPIPostNoodleApplicationsHandler = noodle_api.PostNoodleApplicationsHandlerFunc(func(params noodle_api.PostNoodleApplicationsParams, principal *models.Principal) middleware.Responder {
		return HandlerApplicationPost(db, params, principal)
	})
	api.NoodleAPIDeleteNoodleApplicationsHandler = noodle_api.DeleteNoodleApplicationsHandlerFunc(func(params noodle_api.DeleteNoodleApplicationsParams, principal *models.Principal) middleware.Responder {
		return HandlerApplicationDelete(db, params, principal)
	})

}
