package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/heimdall"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
)

func HandleHeimdallRefresh(db database.Database, heimdall heimdall.Heimdall, params noodle_api.GetNoodleHeimdallReloadParams, principal *models.Principal) middleware.Responder {
	Logger.Info().Msg("Starting Heimdall Refresh")

	err := heimdall.UpdateFromServer()
	if err != nil {
		Logger.Error().Err(err).Msg("heimdall.UpdateFromServer")
		return noodle_api.NewGetNoodleLdapReloadConflict().WithPayload(&models.Error{Message: err.Error()})
	}

	Logger.Info().Msg("Finished Heimdall Refresh")
	return noodle_api.NewGetNoodleHeimdallReloadOK()

}
