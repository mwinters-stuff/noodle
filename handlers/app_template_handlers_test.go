package handlers_test

import (
	"errors"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/mwinters-stuff/noodle/handlers"
	handler_mocks "github.com/mwinters-stuff/noodle/handlers/mocks"
	"github.com/mwinters-stuff/noodle/noodle/database/mocks"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AppTemplateHandlersTestSuite struct {
	suite.Suite
	mockDatabase         *mocks.Database
	mockTables           *mocks.Tables
	mockAppTemplateTable *mocks.AppTemplateTable
}

func (suite *AppTemplateHandlersTestSuite) SetupSuite() {

}

func (suite *AppTemplateHandlersTestSuite) SetupTest() {
	suite.mockDatabase = mocks.NewDatabase(suite.T())
	suite.mockTables = mocks.NewTables(suite.T())
	suite.mockAppTemplateTable = mocks.NewAppTemplateTable(suite.T())

	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(1)
	suite.mockTables.EXPECT().AppTemplateTable().Return(suite.mockAppTemplateTable).Times(1)

}

func (suite *AppTemplateHandlersTestSuite) TearDownTest() {

}
func (suite *AppTemplateHandlersTestSuite) TearSuite() {

}

func (suite *AppTemplateHandlersTestSuite) TestHandlerAppTemplatesGet() {
	suite.mockAppTemplateTable.EXPECT().Search("guard").Return([]*models.ApplicationTemplate{
		{
			Appid:          "140902edbcc424c09736af28ab2de604c3bde936",
			Name:           "AdGuard Home",
			Website:        "https://github.com/AdguardTeam/AdGuardHome",
			License:        "GNU General Public License v3.0 only",
			Description:    "AdGuard Home is a network-wide software for blocking ads.",
			Enhanced:       true,
			TileBackground: "light",
			Icon:           "adguardhome.png",
			SHA:            "ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7",
		},
	}, nil)

	pr := models.Principal("")

	params := noodle_api.NewGetNoodleAppTemplatesParams()
	params.Search = "guard"
	response := handlers.HandlerAppTemplates(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`[{"Appid":"140902edbcc424c09736af28ab2de604c3bde936","Description":"AdGuard Home is a network-wide software for blocking ads.","Enhanced":true,"Icon":"adguardhome.png","License":"GNU General Public License v3.0 only","Name":"AdGuard Home","SHA":"ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7","Website":"https://github.com/AdguardTeam/AdGuardHome","tile_background":"light"}]`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *AppTemplateHandlersTestSuite) TestHandlerAppTemplatesGetError() {
	suite.mockAppTemplateTable.EXPECT().Search("guard").Return([]*models.ApplicationTemplate{}, errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewGetNoodleAppTemplatesParams()
	params.Search = "guard"
	response := handlers.HandlerAppTemplates(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()

	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func TestAppTemplateHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(AppTemplateHandlersTestSuite))
}
