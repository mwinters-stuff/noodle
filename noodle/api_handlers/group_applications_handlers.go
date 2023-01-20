package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
)

func HandlerGroupApplicationGet(db database.Database, params noodle_api.GetNoodleGroupApplicationsParams, principal *models.Principal) middleware.Responder {
	applicationTabs, err := db.Tables().GroupApplicationsTable().GetGroupApps(params.GroupID)
	if err != nil {
		return noodle_api.NewGetNoodleGroupApplicationsConflict().WithPayload(&models.Error{Message: err.Error()})
	}
	return noodle_api.NewGetNoodleGroupApplicationsOK().WithPayload(applicationTabs)
}

func HandlerGroupApplicationPost(db database.Database, params noodle_api.PostNoodleGroupApplicationsParams, principal *models.Principal) middleware.Responder {
	err := db.Tables().GroupApplicationsTable().Insert(params.GroupApplication)

	if err != nil {
		return noodle_api.NewPostNoodleTabsConflict().WithPayload(&models.Error{Message: err.Error()})
	}

	return noodle_api.NewPostNoodleGroupApplicationsOK().WithPayload(params.GroupApplication)
}

func HandlerGroupApplicationDelete(db database.Database, params noodle_api.DeleteNoodleGroupApplicationsParams, principal *models.Principal) middleware.Responder {
	err := db.Tables().GroupApplicationsTable().Delete(params.GroupApplicationID)
	if err != nil {
		return noodle_api.NewDeleteNoodleGroupApplicationsConflict().WithPayload(&models.Error{Message: err.Error()})
	}
	return noodle_api.NewDeleteNoodleGroupApplicationsOK()

}
