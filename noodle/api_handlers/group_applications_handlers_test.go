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

type GroupApplicationsHandlersTestSuite struct {
	suite.Suite
	mockDatabase               *mocks.Database
	mockTables                 *mocks.Tables
	mockGroupApplicationsTable *mocks.GroupApplicationsTable
	api                        *operations.NoodleAPI
}

func (suite *GroupApplicationsHandlersTestSuite) SetupSuite() {
	database.Logger = log.Output(nil)
	ldap_handler.Logger = log.Output(nil)
}

func (suite *GroupApplicationsHandlersTestSuite) SetupTest() {
	suite.mockDatabase = mocks.NewDatabase(suite.T())
	suite.mockTables = mocks.NewTables(suite.T())
	suite.mockGroupApplicationsTable = mocks.NewGroupApplicationsTable(suite.T())

	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Once()
	suite.mockTables.EXPECT().GroupApplicationsTable().Return(suite.mockGroupApplicationsTable).Once()

	suite.api = &operations.NoodleAPI{}
	api_handlers.RegisterGroupApplicationsApiHandlers(suite.api, suite.mockDatabase)

	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleGroupApplicationsHandler)
	require.NotNil(suite.T(), suite.api.NoodleAPIPostNoodleGroupApplicationsHandler)
	require.NotNil(suite.T(), suite.api.NoodleAPIDeleteNoodleGroupApplicationsHandler)

}

func (suite *GroupApplicationsHandlersTestSuite) TearDownTest() {

}
func (suite *GroupApplicationsHandlersTestSuite) TearSuite() {

}

func (suite *GroupApplicationsHandlersTestSuite) TestHandlerGroupApplicationsGet() {
	suite.mockGroupApplicationsTable.EXPECT().GetGroupApps(int64(1)).Return([]*models.GroupApplications{
		{
			ID:            1,
			GroupID:       1,
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
			GroupID:       1,
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

	params := noodle_api.NewGetNoodleGroupApplicationsParams()
	params.GroupID = 1
	response := suite.api.NoodleAPIGetNoodleGroupApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`[{"Application":{"Description":"AdGuard Home is a network-wide software for blocking ads.","Enhanced":true,"Icon":"adguardhome.png","Id":1,"License":"GNU General Public License v3.0 only","Name":"AdGuard Home","TileBackground":"light","Website":"https://github.com/AdguardTeam/AdGuardHome"},"ApplicationId":1,"GroupId":1,"Id":1},{"Application":{"Description":"Adminer.","Icon":"adminer.svg","Id":2,"License":"Apache License 2.0","Name":"Adminer","TileBackground":"light","Website":"https://www.adminer.org"},"ApplicationId":2,"GroupId":1,"Id":2}]`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *GroupApplicationsHandlersTestSuite) TestHandlerGroupApplicationsGetError() {
	suite.mockGroupApplicationsTable.EXPECT().GetGroupApps(int64(1)).Return([]*models.GroupApplications{}, errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewGetNoodleGroupApplicationsParams()
	params.GroupID = 1
	response := suite.api.NoodleAPIGetNoodleGroupApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *GroupApplicationsHandlersTestSuite) TestHandlerGroupApplicationsInsert() {
	suite.mockGroupApplicationsTable.EXPECT().Insert(&models.GroupApplications{
		GroupID:       1,
		ApplicationID: 1,
	}).Return(nil)

	pr := models.Principal("")

	params := noodle_api.NewPostNoodleGroupApplicationsParams()
	params.GroupApplication = &models.GroupApplications{
		GroupID:       1,
		ApplicationID: 1,
	}

	response := suite.api.NoodleAPIPostNoodleGroupApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`{"ApplicationId":1,"GroupId":1}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *GroupApplicationsHandlersTestSuite) TestHandlerGroupApplicationsInsertError() {
	suite.mockGroupApplicationsTable.EXPECT().Insert(&models.GroupApplications{
		GroupID:       1,
		ApplicationID: 1,
	}).Return(errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewPostNoodleGroupApplicationsParams()
	params.GroupApplication = &models.GroupApplications{
		GroupID:       1,
		ApplicationID: 1,
	}

	response := suite.api.NoodleAPIPostNoodleGroupApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *GroupApplicationsHandlersTestSuite) TestHandlerGroupApplicationsDelete() {
	suite.mockGroupApplicationsTable.EXPECT().Delete(int64(1)).Return(nil)

	pr := models.Principal("")

	params := noodle_api.NewDeleteNoodleGroupApplicationsParams()
	params.GroupApplicationID = 1

	response := suite.api.NoodleAPIDeleteNoodleGroupApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().Header().Once().Return(nil)
	mockWriter.EXPECT().WriteHeader(200).Once()

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *GroupApplicationsHandlersTestSuite) TestHandlerGroupApplicationsDeleteError() {
	suite.mockGroupApplicationsTable.EXPECT().Delete(int64(1)).Return(errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewDeleteNoodleGroupApplicationsParams()
	params.GroupApplicationID = 1

	response := suite.api.NoodleAPIDeleteNoodleGroupApplicationsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)
	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func TestGroupApplicationsHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(GroupApplicationsHandlersTestSuite))
}
