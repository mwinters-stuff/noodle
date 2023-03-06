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

type UserApplicationsHandlersTestSuite struct {
	suite.Suite
	mockDatabase              *mocks.Database
	mockTables                *mocks.Tables
	mockUserApplicationsTable *mocks.UserApplicationsTable
	api                       *operations.NoodleAPI
}

func (suite *UserApplicationsHandlersTestSuite) SetupSuite() {
	database.Logger = log.Output(nil)
	ldap_handler.Logger = log.Output(nil)
}

func (suite *UserApplicationsHandlersTestSuite) SetupTest() {
	suite.mockDatabase = mocks.NewDatabase(suite.T())
	suite.mockTables = mocks.NewTables(suite.T())
	suite.mockUserApplicationsTable = mocks.NewUserApplicationsTable(suite.T())

	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Once()
	suite.mockTables.EXPECT().UserApplicationsTable().Return(suite.mockUserApplicationsTable).Once()

	suite.api = &operations.NoodleAPI{}
	api_handlers.RegisterUserApplicationsApiHandlers(suite.api, suite.mockDatabase)

	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleUserApplicationsHandler)
	require.NotNil(suite.T(), suite.api.NoodleAPIPostNoodleUserApplicationsHandler)
	require.NotNil(suite.T(), suite.api.NoodleAPIDeleteNoodleUserApplicationsHandler)
	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleUserAllowedApplicationsHandler)
}

func (suite *UserApplicationsHandlersTestSuite) TearDownTest() {

}
func (suite *UserApplicationsHandlersTestSuite) TearDownSuite() {

}

func (suite *UserApplicationsHandlersTestSuite) TestHandlerUserApplicationsGet() {
	suite.mockUserApplicationsTable.EXPECT().GetUserApps(int64(1)).Return([]*models.UserApplications{
		{
			ID:            1,
			UserID:        1,
			ApplicationID: 1,
			Application: &models.Application{
				ID:             1,
				TemplateAppid:  "",
				Name:           "AdGuard Home",
				Website:        "https://github.com/AdguardTeam/AdGuardHome",
				License:        "GNU General Public License v3.0 only",
				Description:    "AdGuard Home is a network-wide software for blocking ads.",
				Enhanced:       true,
				TileBackground: "light",
				TextColor:      "dark",
				Icon:           "adguardhome.png",
			},
		},
		{
			ID:            2,
			UserID:        1,
			ApplicationID: 2,
			Application: &models.Application{
				ID:             2,
				Name:           "Adminer",
				TemplateAppid:  "",
				Website:        "https://www.adminer.org",
				License:        "Apache License 2.0",
				Description:    "Adminer.",
				Enhanced:       false,
				TileBackground: "light",
				TextColor:      "dark",
				Icon:           "adminer.svg",
			},
		},
	}, nil)

	pr := models.Principal("")

	params := noodle_api.NewGetNoodleUserApplicationsParams()
	params.UserID = 1
	response := suite.api.NoodleAPIGetNoodleUserApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`[{"Application":{"Description":"AdGuard Home is a network-wide software for blocking ads.","Enhanced":true,"Icon":"adguardhome.png","Id":1,"License":"GNU General Public License v3.0 only","Name":"AdGuard Home","TextColor":"dark","TileBackground":"light","Website":"https://github.com/AdguardTeam/AdGuardHome"},"ApplicationId":1,"Id":1,"UserId":1},{"Application":{"Description":"Adminer.","Icon":"adminer.svg","Id":2,"License":"Apache License 2.0","Name":"Adminer","TextColor":"dark","TileBackground":"light","Website":"https://www.adminer.org"},"ApplicationId":2,"Id":2,"UserId":1}]`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *UserApplicationsHandlersTestSuite) TestHandlerUserApplicationsGetError() {
	suite.mockUserApplicationsTable.EXPECT().GetUserApps(int64(1)).Return([]*models.UserApplications{}, errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewGetNoodleUserApplicationsParams()
	params.UserID = 1
	response := suite.api.NoodleAPIGetNoodleUserApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *UserApplicationsHandlersTestSuite) TestHandlerUserApplicationsInsert() {
	suite.mockUserApplicationsTable.EXPECT().Insert(&models.UserApplications{
		UserID:        1,
		ApplicationID: 1,
	}).Return(nil)

	pr := models.Principal("")

	params := noodle_api.NewPostNoodleUserApplicationsParams()
	params.UserApplication = &models.UserApplications{
		UserID:        1,
		ApplicationID: 1,
	}

	response := suite.api.NoodleAPIPostNoodleUserApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`{"ApplicationId":1,"UserId":1}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *UserApplicationsHandlersTestSuite) TestHandlerUserApplicationsInsertError() {
	suite.mockUserApplicationsTable.EXPECT().Insert(&models.UserApplications{
		UserID:        1,
		ApplicationID: 1,
	}).Return(errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewPostNoodleUserApplicationsParams()
	params.UserApplication = &models.UserApplications{
		UserID:        1,
		ApplicationID: 1,
	}

	response := suite.api.NoodleAPIPostNoodleUserApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *UserApplicationsHandlersTestSuite) TestHandlerUserApplicationsDelete() {
	suite.mockUserApplicationsTable.EXPECT().Delete(int64(1)).Return(nil)

	pr := models.Principal("")

	params := noodle_api.NewDeleteNoodleUserApplicationsParams()
	params.UserApplicationID = 1

	response := suite.api.NoodleAPIDeleteNoodleUserApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().Header().Once().Return(nil)
	mockWriter.EXPECT().WriteHeader(200).Once()

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *UserApplicationsHandlersTestSuite) TestHandlerUserApplicationsDeleteError() {
	suite.mockUserApplicationsTable.EXPECT().Delete(int64(1)).Return(errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewDeleteNoodleUserApplicationsParams()
	params.UserApplicationID = 1

	response := suite.api.NoodleAPIDeleteNoodleUserApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)
	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *UserApplicationsHandlersTestSuite) TestHandlerUserAllowedApplicationsGet() {
	suite.mockUserApplicationsTable.EXPECT().GetUserAllowdApplications(int64(1)).Return([]*models.UsersApplicationItem{
		{
			Application: &models.Application{
				Description:    "application_tab_1",
				Enhanced:       false,
				Icon:           "string",
				ID:             2,
				License:        "string",
				Name:           "applicationtab1",
				TemplateAppid:  "",
				TileBackground: "string",
				TextColor:      "dark",
				Website:        "string",
			},
			DisplayOrder: 0,
			TabID:        1,
		},
		{
			Application: &models.Application{
				Description:    "user custom app",
				Enhanced:       false,
				Icon:           "string",
				ID:             1,
				License:        "string",
				Name:           "usercustomapp",
				TemplateAppid:  "",
				TileBackground: "string",
				TextColor:      "dark",
				Website:        "string",
			},
			DisplayOrder: 6,
			TabID:        1,
		},
	}, nil)

	pr := models.Principal("")

	params := noodle_api.NewGetNoodleUserAllowedApplicationsParams()
	params.UserID = 1
	response := suite.api.NoodleAPIGetNoodleUserAllowedApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`[{"Application":{"Description":"application_tab_1","Icon":"string","Id":2,"License":"string","Name":"applicationtab1","TextColor":"dark","TileBackground":"string","Website":"string"},"TabId":1},{"Application":{"Description":"user custom app","Icon":"string","Id":1,"License":"string","Name":"usercustomapp","TextColor":"dark","TileBackground":"string","Website":"string"},"DisplayOrder":6,"TabId":1}]`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *UserApplicationsHandlersTestSuite) TestHandlerUserAllowedApplicationsGetError() {
	suite.mockUserApplicationsTable.EXPECT().GetUserAllowdApplications(int64(1)).Return([]*models.UsersApplicationItem{}, errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewGetNoodleUserAllowedApplicationsParams()
	params.UserID = 1
	response := suite.api.NoodleAPIGetNoodleUserAllowedApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func TestUserApplicationsHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(UserApplicationsHandlersTestSuite))
}
