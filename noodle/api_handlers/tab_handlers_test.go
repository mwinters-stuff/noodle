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

type TabHandlersTestSuite struct {
	suite.Suite
	mockDatabase *mocks.Database
	mockTables   *mocks.Tables
	mockTabTable *mocks.TabTable
}

func (suite *TabHandlersTestSuite) SetupSuite() {

}

func (suite *TabHandlersTestSuite) SetupTest() {
	suite.mockDatabase = mocks.NewDatabase(suite.T())
	suite.mockTables = mocks.NewTables(suite.T())
	suite.mockTabTable = mocks.NewTabTable(suite.T())

	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Once()
	suite.mockTables.EXPECT().TabTable().Return(suite.mockTabTable).Once()

}

func (suite *TabHandlersTestSuite) TearDownTest() {

}
func (suite *TabHandlersTestSuite) TearSuite() {

}

func (suite *TabHandlersTestSuite) TestHandlerTabsGet() {
	suite.mockTabTable.EXPECT().GetAll().Return([]*models.Tab{
		{
			ID:           1,
			Label:        "Servers",
			DisplayOrder: 1,
		},
		{
			ID:           2,
			Label:        "Apps",
			DisplayOrder: 2,
		},
	}, nil)

	pr := models.Principal("")

	response := handlers.HandlerTabGet(suite.mockDatabase, noodle_api.NewGetNoodleTabsParams(), &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`[{"DisplayOrder":1,"Id":1,"Label":"Servers"},{"DisplayOrder":2,"Id":2,"Label":"Apps"}]`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *TabHandlersTestSuite) TestHandlerTabsGetError() {
	suite.mockTabTable.EXPECT().GetAll().Return([]*models.Tab{}, errors.New("failed"))

	pr := models.Principal("")

	response := handlers.HandlerTabGet(suite.mockDatabase, noodle_api.NewGetNoodleTabsParams(), &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *TabHandlersTestSuite) TestHandlerTabsInsert() {
	suite.mockTabTable.EXPECT().Insert(&models.Tab{
		Label:        "Servers",
		DisplayOrder: 1,
	}).Return(nil)

	pr := models.Principal("")

	params := noodle_api.NewPostNoodleTabsParams()
	params.Tab = &models.Tab{
		Label:        "Servers",
		DisplayOrder: 1,
	}
	params.Action = "insert"

	response := handlers.HandlerTabPost(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`{"DisplayOrder":1,"Label":"Servers"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *TabHandlersTestSuite) TestHandlerTabsInsertError() {
	suite.mockTabTable.EXPECT().Insert(&models.Tab{
		Label:        "Servers",
		DisplayOrder: 1,
	}).Return(errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewPostNoodleTabsParams()
	params.Tab = &models.Tab{
		Label:        "Servers",
		DisplayOrder: 1,
	}
	params.Action = "insert"

	response := handlers.HandlerTabPost(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *TabHandlersTestSuite) TestHandlerTabsUpdate() {
	suite.mockTabTable.EXPECT().Update(models.Tab{
		ID:           1,
		Label:        "Servers",
		DisplayOrder: 1,
	}).Return(nil)

	pr := models.Principal("")

	params := noodle_api.NewPostNoodleTabsParams()
	params.Tab = &models.Tab{
		ID:           1,
		Label:        "Servers",
		DisplayOrder: 1,
	}
	params.Action = "update"

	response := handlers.HandlerTabPost(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`{"DisplayOrder":1,"Id":1,"Label":"Servers"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *TabHandlersTestSuite) TestHandlerTabsUpdateError() {
	suite.mockTabTable.EXPECT().Update(models.Tab{
		ID:           1,
		Label:        "Servers",
		DisplayOrder: 1,
	}).Return(errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewPostNoodleTabsParams()
	params.Tab = &models.Tab{
		ID:           1,
		Label:        "Servers",
		DisplayOrder: 1,
	}
	params.Action = "update"

	response := handlers.HandlerTabPost(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)
	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *TabHandlersTestSuite) TestHandlerTabsDelete() {
	suite.mockTabTable.EXPECT().Delete(int64(1)).Return(nil)

	pr := models.Principal("")

	params := noodle_api.NewDeleteNoodleTabsParams()
	params.Tabid = 1

	response := handlers.HandlerTabDelete(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().Header().Once().Return(nil)
	mockWriter.EXPECT().WriteHeader(200).Once()

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *TabHandlersTestSuite) TestHandlerTabsDeleteError() {
	suite.mockTabTable.EXPECT().Delete(int64(1)).Return(errors.New("failed"))

	pr := models.Principal("")

	params := noodle_api.NewDeleteNoodleTabsParams()
	params.Tabid = 1

	response := handlers.HandlerTabDelete(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)
	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func TestTabHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(TabHandlersTestSuite))
}
