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

}

func (suite *UserApplicationsHandlersTestSuite) TearDownTest() {

}
func (suite *UserApplicationsHandlersTestSuite) TearSuite() {

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
				Icon:           "adminer.svg",
			},
		},
	}, nil)

	pr := models.Principal("")

	params := noodle_api.NewGetNoodleUserApplicationsParams()
	params.UserID = 1
	response := api_handlers.HandlerUserApplicationGet(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`[{"Application":{"Description":"AdGuard Home is a network-wide software for blocking ads.","Enhanced":true,"Icon":"adguardhome.png","Id":1,"License":"GNU General Public License v3.0 only","Name":"AdGuard Home","TileBackground":"light","Website":"https://github.com/AdguardTeam/AdGuardHome"},"ApplicationId":1,"Id":1,"UserId":1},{"Application":{"Description":"Adminer.","Icon":"adminer.svg","Id":2,"License":"Apache License 2.0","Name":"Adminer","TileBackground":"light","Website":"https://www.adminer.org"},"ApplicationId":2,"Id":2,"UserId":1}]`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *UserApplicationsHandlersTestSuite) TestHandlerUserApplicationsGetError() {
	suite.mockUserApplicationsTable.EXPECT().GetUserApps(int64(1)).Return([]*models.UserApplications{}, errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewGetNoodleUserApplicationsParams()
	params.UserID = 1
	response := api_handlers.HandlerUserApplicationGet(suite.mockDatabase, params, &pr)
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

	response := api_handlers.HandlerUserApplicationPost(suite.mockDatabase, params, &pr)
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

	response := api_handlers.HandlerUserApplicationPost(suite.mockDatabase, params, &pr)
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

	response := api_handlers.HandlerUserApplicationDelete(suite.mockDatabase, params, &pr)
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

	response := api_handlers.HandlerUserApplicationDelete(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)
	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func TestUserApplicationsHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(UserApplicationsHandlersTestSuite))
}
