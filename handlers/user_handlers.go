package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
)

func HandlerUsers(db database.Database, params noodle_api.GetNoodleUsersParams, principal *models.Principal) middleware.Responder {
	var users []*models.User
	var err error
	var user models.User
	if params.Userid != nil && *params.Userid > int64(0) {
		user, err = db.Tables().UserTable().GetID(*params.Userid)
		users = append(users, &user)
	} else {
		users, err = db.Tables().UserTable().GetAll()
	}
	if err != nil {
		return noodle_api.NewGetNoodleUsersConflict().WithPayload(&models.Error{Message: err.Error()})
	}
	return noodle_api.NewGetNoodleUsersOK().WithPayload(users)
}
