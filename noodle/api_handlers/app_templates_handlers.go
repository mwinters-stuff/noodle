package api_handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
)

func HandlerAppTemplates(db database.Database, params noodle_api.GetNoodleAppTemplatesParams, principal *models.Principal) middleware.Responder {
	templates, err := db.Tables().AppTemplateTable().Search(params.Search)

	if err != nil {
		return noodle_api.NewGetNoodleAppTemplatesConflict().WithPayload(&models.Error{Message: err.Error()})
	}
	return noodle_api.NewGetNoodleAppTemplatesOK().WithPayload(templates)
}
