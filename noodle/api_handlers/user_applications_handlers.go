package api_handlers

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
)

func HandlerUserApplicationGet(db database.Database, params noodle_api.GetNoodleUserApplicationsParams, principal *models.Principal) middleware.Responder {
	applicationTabs, err := db.Tables().UserApplicationsTable().GetUserApps(params.UserID)
	if err != nil {
		return noodle_api.NewGetNoodleUserApplicationsConflict().WithPayload(&models.Error{Message: err.Error()})
	}
	return noodle_api.NewGetNoodleUserApplicationsOK().WithPayload(applicationTabs)
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
