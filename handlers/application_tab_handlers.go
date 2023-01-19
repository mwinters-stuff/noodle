package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
)

func HandlerApplicationTabGet(db database.Database, params noodle_api.GetNoodleApplicationTabsParams, principal *models.Principal) middleware.Responder {
	applicationTabs, err := db.Tables().ApplicationTabTable().GetTabApps(params.TabID)
	if err != nil {
		return noodle_api.NewGetNoodleApplicationTabsConflict().WithPayload(&models.Error{Message: err.Error()})
	}
	return noodle_api.NewGetNoodleApplicationTabsOK().WithPayload(applicationTabs)
}

func HandlerApplicationTabPost(db database.Database, params noodle_api.PostNoodleApplicationTabsParams, principal *models.Principal) middleware.Responder {
	if params.Action == "insert" {
		err := db.Tables().ApplicationTabTable().Insert(params.ApplicationTab)

		if err != nil {
			return noodle_api.NewPostNoodleTabsConflict().WithPayload(&models.Error{Message: err.Error()})
		}
	}
	if params.Action == "update" {
		err := db.Tables().ApplicationTabTable().Update(*params.ApplicationTab)

		if err != nil {
			return noodle_api.NewPostNoodleTabsConflict().WithPayload(&models.Error{Message: err.Error()})
		}
	}

	return noodle_api.NewPostNoodleApplicationTabsOK().WithPayload(params.ApplicationTab)
}

func HandlerApplicationTabDelete(db database.Database, params noodle_api.DeleteNoodleApplicationTabsParams, principal *models.Principal) middleware.Responder {
	err := db.Tables().ApplicationTabTable().Delete(params.ApplicationTabID)
	if err != nil {
		return noodle_api.NewDeleteNoodleApplicationTabsConflict().WithPayload(&models.Error{Message: err.Error()})
	}
	return noodle_api.NewDeleteNoodleApplicationTabsOK()

}
