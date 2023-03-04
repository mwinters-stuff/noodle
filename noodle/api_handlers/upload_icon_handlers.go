package api_handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mwinters-stuff/noodle/noodle/heimdall"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
)

var (
	RegisterUploadIconApiHandlers = RegisterUploadIconApiHandlersImpl
)

func HandlerUploadIcons(heimdall heimdall.Heimdall, params noodle_api.PostNoodleUploadIconParams, principal *models.Principal) middleware.Responder {
	filename := params.HTTPRequest.MultipartForm.File["upfile"][0].Filename
	err := heimdall.UploadIcon(filename, params.Upfile)
	if err != nil {
		return noodle_api.NewPostNoodleUploadIconConflict().WithPayload(&models.Error{Message: err.Error()})
	}

	return noodle_api.NewPostNoodleUploadIconOK()
}

func HandlerGetIcons(heimdall heimdall.Heimdall, params noodle_api.GetNoodleUploadIconParams, principal *models.Principal) middleware.Responder {
	filenames, err := heimdall.ListIcons()
	if err != nil {
		return noodle_api.NewGetNoodleUploadIconConflict().WithPayload(&models.Error{Message: err.Error()})
	}

	return noodle_api.NewGetNoodleUploadIconOK().WithPayload(filenames)

}

func RegisterUploadIconApiHandlersImpl(api *operations.NoodleAPI, heimdall heimdall.Heimdall) {
	api.NoodleAPIPostNoodleUploadIconHandler = noodle_api.PostNoodleUploadIconHandlerFunc(func(params noodle_api.PostNoodleUploadIconParams, principal *models.Principal) middleware.Responder {
		return HandlerUploadIcons(heimdall, params, principal)
	})

	api.NoodleAPIGetNoodleUploadIconHandler = noodle_api.GetNoodleUploadIconHandlerFunc(func(params noodle_api.GetNoodleUploadIconParams, principal *models.Principal) middleware.Responder {
		return HandlerGetIcons(heimdall, params, principal)
	})
}
