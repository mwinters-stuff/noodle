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

type ApplicationTabHandlersTestSuite struct {
	suite.Suite
	mockDatabase            *mocks.Database
	mockTables              *mocks.Tables
	mockApplicationTabTable *mocks.ApplicationTabTable
}

func (suite *ApplicationTabHandlersTestSuite) SetupSuite() {
	database.Logger = log.Output(nil)
	ldap_handler.Logger = log.Output(nil)
}

func (suite *ApplicationTabHandlersTestSuite) SetupTest() {
	suite.mockDatabase = mocks.NewDatabase(suite.T())
	suite.mockTables = mocks.NewTables(suite.T())
	suite.mockApplicationTabTable = mocks.NewApplicationTabTable(suite.T())

	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Once()
	suite.mockTables.EXPECT().ApplicationTabTable().Return(suite.mockApplicationTabTable).Once()

}

func (suite *ApplicationTabHandlersTestSuite) TearDownTest() {

}
func (suite *ApplicationTabHandlersTestSuite) TearSuite() {

}

func (suite *ApplicationTabHandlersTestSuite) TestHandlerApplicationTabsGet() {
	suite.mockApplicationTabTable.EXPECT().GetTabApps(int64(1)).Return([]*models.ApplicationTab{
		{
			ID:            1,
			TabID:         1,
			ApplicationID: 1,
			DisplayOrder:  1,
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
			TabID:         1,
			ApplicationID: 2,
			DisplayOrder:  2,
			Application: &models.Application{
				ID:             2,
				Name:           "Adminer",
				TemplateAppid:  "",
				Website:        "https://www.adminer.org",
				License:        "Apache License 2.0",
				Description:    "Adminer",
				Enhanced:       false,
				TileBackground: "light",
				Icon:           "adminer.svg",
			},
		},
	}, nil)

	pr := models.Principal("")

	params := noodle_api.NewGetNoodleApplicationTabsParams()
	params.TabID = 1
	response := api_handlers.HandlerApplicationTabGet(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`[{"Application":{"Description":"AdGuard Home is a network-wide software for blocking ads.","Enhanced":true,"Icon":"adguardhome.png","Id":1,"License":"GNU General Public License v3.0 only","Name":"AdGuard Home","TileBackground":"light","Website":"https://github.com/AdguardTeam/AdGuardHome"},"ApplicationId":1,"DisplayOrder":1,"Id":1,"TabId":1},{"Application":{"Description":"Adminer","Icon":"adminer.svg","Id":2,"License":"Apache License 2.0","Name":"Adminer","TileBackground":"light","Website":"https://www.adminer.org"},"ApplicationId":2,"DisplayOrder":2,"Id":2,"TabId":1}]`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *ApplicationTabHandlersTestSuite) TestHandlerApplicationTabsGetError() {
	suite.mockApplicationTabTable.EXPECT().GetTabApps(int64(1)).Return([]*models.ApplicationTab{}, errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewGetNoodleApplicationTabsParams()
	params.TabID = 1
	response := api_handlers.HandlerApplicationTabGet(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *ApplicationTabHandlersTestSuite) TestHandlerApplicationTabsInsert() {
	suite.mockApplicationTabTable.EXPECT().Insert(&models.ApplicationTab{
		TabID:         1,
		ApplicationID: 1,
		DisplayOrder:  1,
	}).Return(nil)

	pr := models.Principal("")

	params := noodle_api.NewPostNoodleApplicationTabsParams()
	params.ApplicationTab = &models.ApplicationTab{
		TabID:         1,
		ApplicationID: 1,
		DisplayOrder:  1,
	}
	params.Action = "insert"

	response := api_handlers.HandlerApplicationTabPost(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`{"ApplicationId":1,"DisplayOrder":1,"TabId":1}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *ApplicationTabHandlersTestSuite) TestHandlerApplicationTabsInsertError() {
	suite.mockApplicationTabTable.EXPECT().Insert(&models.ApplicationTab{
		TabID:         1,
		ApplicationID: 1,
		DisplayOrder:  1,
	}).Return(errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewPostNoodleApplicationTabsParams()
	params.ApplicationTab = &models.ApplicationTab{
		TabID:         1,
		ApplicationID: 1,
		DisplayOrder:  1,
	}
	params.Action = "insert"

	response := api_handlers.HandlerApplicationTabPost(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *ApplicationTabHandlersTestSuite) TestHandlerApplicationTabsUpdate() {
	suite.mockApplicationTabTable.EXPECT().Update(models.ApplicationTab{
		ID:            1,
		TabID:         1,
		ApplicationID: 1,
		DisplayOrder:  2,
	}).Return(nil)

	pr := models.Principal("")

	params := noodle_api.NewPostNoodleApplicationTabsParams()
	params.ApplicationTab = &models.ApplicationTab{
		ID:            1,
		TabID:         1,
		ApplicationID: 1,
		DisplayOrder:  2,
	}
	params.Action = "update"

	response := api_handlers.HandlerApplicationTabPost(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`{"ApplicationId":1,"DisplayOrder":2,"Id":1,"TabId":1}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *ApplicationTabHandlersTestSuite) TestHandlerApplicationTabsUpdateError() {
	suite.mockApplicationTabTable.EXPECT().Update(models.ApplicationTab{
		ID:            1,
		TabID:         1,
		ApplicationID: 1,
		DisplayOrder:  2,
	}).Return(errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewPostNoodleApplicationTabsParams()
	params.ApplicationTab = &models.ApplicationTab{
		ID:            1,
		TabID:         1,
		ApplicationID: 1,
		DisplayOrder:  2,
	}
	params.Action = "update"

	response := api_handlers.HandlerApplicationTabPost(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)
	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *ApplicationTabHandlersTestSuite) TestHandlerApplicationTabsDelete() {
	suite.mockApplicationTabTable.EXPECT().Delete(int64(1)).Return(nil)

	pr := models.Principal("")

	params := noodle_api.NewDeleteNoodleApplicationTabsParams()
	params.ApplicationTabID = 1

	response := api_handlers.HandlerApplicationTabDelete(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().Header().Once().Return(nil)
	mockWriter.EXPECT().WriteHeader(200).Once()

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *ApplicationTabHandlersTestSuite) TestHandlerApplicationTabsDeleteError() {
	suite.mockApplicationTabTable.EXPECT().Delete(int64(1)).Return(errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewDeleteNoodleApplicationTabsParams()
	params.ApplicationTabID = 1

	response := api_handlers.HandlerApplicationTabDelete(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)
	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func TestApplicationTabHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(ApplicationTabHandlersTestSuite))
}
