package api_handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
)

func HandlerGroups(db database.Database, params noodle_api.GetNoodleGroupsParams, principal *models.Principal) middleware.Responder {
	var groups []*models.Group
	var err error
	var group models.Group
	if params.Groupid != nil && *params.Groupid > int64(0) {
		group, err = db.Tables().GroupTable().GetID(*params.Groupid)
		groups = append(groups, &group)
	} else {
		groups, err = db.Tables().GroupTable().GetAll()
	}
	if err != nil {
		return noodle_api.NewGetNoodleGroupsConflict().WithPayload(&models.Error{Message: err.Error()})
	}
	return noodle_api.NewGetNoodleGroupsOK().WithPayload(groups)
}
