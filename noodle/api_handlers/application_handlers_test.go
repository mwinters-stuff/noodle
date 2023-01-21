package api_handlers_test

import (
	"errors"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/mwinters-stuff/noodle/noodle/api_handlers"
	handler_mocks "github.com/mwinters-stuff/noodle/noodle/api_handlers/mocks"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/database/mocks"
	"github.com/mwinters-stuff/noodle/noodle/ldap_handler"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ApplicationsHandlersTestSuite struct {
	suite.Suite
	mockDatabase          *mocks.Database
	mockTables            *mocks.Tables
	mockApplicationsTable *mocks.ApplicationsTable
	api                   *operations.NoodleAPI
}

func (suite *ApplicationsHandlersTestSuite) SetupSuite() {
	database.Logger = log.Output(nil)
	ldap_handler.Logger = log.Output(nil)
}

func (suite *ApplicationsHandlersTestSuite) SetupTest() {
	suite.mockDatabase = mocks.NewDatabase(suite.T())
	suite.mockTables = mocks.NewTables(suite.T())
	suite.mockApplicationsTable = mocks.NewApplicationsTable(suite.T())

	suite.api = &operations.NoodleAPI{}
	api_handlers.RegisterApplicationsApiHandlers(suite.api, suite.mockDatabase)

	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleApplicationsHandler)
	require.NotNil(suite.T(), suite.api.NoodleAPIPostNoodleApplicationsHandler)
	require.NotNil(suite.T(), suite.api.NoodleAPIDeleteNoodleApplicationsHandler)

}

func (suite *ApplicationsHandlersTestSuite) TearDownTest() {

}
func (suite *ApplicationsHandlersTestSuite) TearSuite() {

}

func (suite *ApplicationsHandlersTestSuite) TestHandlerApplicationsGetID() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Once()
	suite.mockTables.EXPECT().ApplicationsTable().Return(suite.mockApplicationsTable).Once()

	suite.mockApplicationsTable.EXPECT().GetID(int64(1)).Return(models.Application{
		ID:             1,
		TemplateAppid:  "",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
	}, nil)

	pr := models.Principal("")

	params := noodle_api.NewGetNoodleApplicationsParams()
	var id = int64(1)
	params.ApplicationID = &id
	response := suite.api.NoodleAPIGetNoodleApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`[{"Description":"AdGuard Home is a network-wide software for blocking ads.","Enhanced":true,"Icon":"adguardhome.png","Id":1,"License":"GNU General Public License v3.0 only","Name":"AdGuard Home","TileBackground":"light","Website":"https://github.com/AdguardTeam/AdGuardHome"}]`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *ApplicationsHandlersTestSuite) TestHandlerApplicationsGetTemplate() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Once()
	suite.mockTables.EXPECT().ApplicationsTable().Return(suite.mockApplicationsTable).Once()

	suite.mockApplicationsTable.EXPECT().GetTemplateID("ABCDEF").Return([]*models.Application{
		{
			ID:             1,
			TemplateAppid:  "ABCDEF",
			Name:           "AdGuard Home",
			Website:        "https://github.com/AdguardTeam/AdGuardHome",
			License:        "GNU General Public License v3.0 only",
			Description:    "AdGuard Home is a network-wide software for blocking ads.",
			Enhanced:       true,
			TileBackground: "light",
			Icon:           "adguardhome.png",
		},
	}, nil)

	pr := models.Principal("")

	params := noodle_api.NewGetNoodleApplicationsParams()
	var template = "ABCDEF"
	params.ApplicationTemplate = &template

	response := suite.api.NoodleAPIGetNoodleApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`[{"Description":"AdGuard Home is a network-wide software for blocking ads.","Enhanced":true,"Icon":"adguardhome.png","Id":1,"License":"GNU General Public License v3.0 only","Name":"AdGuard Home","TemplateAppid":"ABCDEF","TileBackground":"light","Website":"https://github.com/AdguardTeam/AdGuardHome"}]`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *ApplicationsHandlersTestSuite) TestHandlerApplicationsGetParamsNullError() {

	pr := models.Principal("")

	params := noodle_api.NewGetNoodleApplicationsParams()

	response := suite.api.NoodleAPIGetNoodleApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"incorrect parameters - both nil"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *ApplicationsHandlersTestSuite) TestHandlerApplicationsGetParamsBothSetError() {

	pr := models.Principal("")

	params := noodle_api.NewGetNoodleApplicationsParams()
	var id = int64(1)
	var template = "abcdef"
	params.ApplicationID = &id
	params.ApplicationTemplate = &template

	response := suite.api.NoodleAPIGetNoodleApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"incorrect parameters - both supplied"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *ApplicationsHandlersTestSuite) TestHandlerApplicationsGetIDError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Once()
	suite.mockTables.EXPECT().ApplicationsTable().Return(suite.mockApplicationsTable).Once()
	suite.mockApplicationsTable.EXPECT().GetID(int64(1)).Return(models.Application{}, errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewGetNoodleApplicationsParams()
	var id = int64(1)
	params.ApplicationID = &id
	response := suite.api.NoodleAPIGetNoodleApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *ApplicationsHandlersTestSuite) TestHandlerApplicationsGetTemplateError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Once()
	suite.mockTables.EXPECT().ApplicationsTable().Return(suite.mockApplicationsTable).Once()
	suite.mockApplicationsTable.EXPECT().GetTemplateID("ABCDEF").Return([]*models.Application{}, errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewGetNoodleApplicationsParams()
	var template = "ABCDEF"
	params.ApplicationTemplate = &template

	response := suite.api.NoodleAPIGetNoodleApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *ApplicationsHandlersTestSuite) TestHandlerApplicationsInsert() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Once()
	suite.mockTables.EXPECT().ApplicationsTable().Return(suite.mockApplicationsTable).Once()
	suite.mockApplicationsTable.EXPECT().Insert(&models.Application{
		TemplateAppid:  "",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
	}).Return(nil)

	pr := models.Principal("")

	params := noodle_api.NewPostNoodleApplicationsParams()
	params.Application = &models.Application{
		TemplateAppid:  "",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
	}

	response := suite.api.NoodleAPIPostNoodleApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`{"Description":"AdGuard Home is a network-wide software for blocking ads.","Enhanced":true,"Icon":"adguardhome.png","License":"GNU General Public License v3.0 only","Name":"AdGuard Home","TileBackground":"light","Website":"https://github.com/AdguardTeam/AdGuardHome"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *ApplicationsHandlersTestSuite) TestHandlerApplicationsInsertError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Once()
	suite.mockTables.EXPECT().ApplicationsTable().Return(suite.mockApplicationsTable).Once()
	suite.mockApplicationsTable.EXPECT().Insert(&models.Application{
		TemplateAppid:  "",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
	}).Return(errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewPostNoodleApplicationsParams()
	params.Application = &models.Application{
		TemplateAppid:  "",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
	}

	response := suite.api.NoodleAPIPostNoodleApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *ApplicationsHandlersTestSuite) TestHandlerApplicationsDelete() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Once()
	suite.mockTables.EXPECT().ApplicationsTable().Return(suite.mockApplicationsTable).Once()
	suite.mockApplicationsTable.EXPECT().Delete(int64(1)).Return(nil)

	pr := models.Principal("")

	params := noodle_api.NewDeleteNoodleApplicationsParams()
	params.ApplicationID = 1

	response := suite.api.NoodleAPIDeleteNoodleApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().Header().Once().Return(nil)
	mockWriter.EXPECT().WriteHeader(200).Once()

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *ApplicationsHandlersTestSuite) TestHandlerApplicationsDeleteError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Once()
	suite.mockTables.EXPECT().ApplicationsTable().Return(suite.mockApplicationsTable).Once()
	suite.mockApplicationsTable.EXPECT().Delete(int64(1)).Return(errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewDeleteNoodleApplicationsParams()
	params.ApplicationID = 1

	response := suite.api.NoodleAPIDeleteNoodleApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)
	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func TestApplicationsHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(ApplicationsHandlersTestSuite))
}
