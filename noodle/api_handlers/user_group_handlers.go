package api_handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
)

func HandlerUserGroups(db database.Database, params noodle_api.GetNoodleUserGroupsParams, principal *models.Principal) middleware.Responder {
	var usergroups []*models.UserGroup
	var err error
	if params.Groupid != nil && *params.Groupid > int64(0) {
		usergroups, err = db.Tables().UserGroupsTable().GetGroup(*params.Groupid)
	} else if params.Userid != nil && *params.Userid > int64(0) {
		usergroups, err = db.Tables().UserGroupsTable().GetUser(*params.Userid)
	} else {
		return noodle_api.NewGetNoodleUserGroupsConflict().WithPayload(&models.Error{Message: "no groupid or userid parameter"})
	}
	if err != nil {
		return noodle_api.NewGetNoodleUserGroupsConflict().WithPayload(&models.Error{Message: err.Error()})
	}
	return noodle_api.NewGetNoodleUserGroupsOK().WithPayload(usergroups)
}
