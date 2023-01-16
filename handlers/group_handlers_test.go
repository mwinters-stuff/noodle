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

type GroupHandlersTestSuite struct {
	suite.Suite
	mockDatabase   *mocks.Database
	mockTables     *mocks.Tables
	mockGroupTable *mocks.GroupTable
}

func (suite *GroupHandlersTestSuite) SetupSuite() {

}

func (suite *GroupHandlersTestSuite) SetupTest() {
	suite.mockDatabase = mocks.NewDatabase(suite.T())
	suite.mockTables = mocks.NewTables(suite.T())
	suite.mockGroupTable = mocks.NewGroupTable(suite.T())

	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(1)
	suite.mockTables.EXPECT().GroupTable().Return(suite.mockGroupTable).Times(1)

}

func (suite *GroupHandlersTestSuite) TearDownTest() {

}
func (suite *GroupHandlersTestSuite) TearSuite() {

}

func (suite *GroupHandlersTestSuite) TestHandlerGroupsGetAll() {
	suite.mockGroupTable.EXPECT().GetAll().Return([]*models.Group{
		{
			ID:   1,
			DN:   "cn=admins,ou=groups,dc=example,dc=nz",
			Name: "Admins",
		},
		{
			ID:   2,
			DN:   "cn=users,ou=groups,dc=example,dc=nz",
			Name: "Users",
		},
	}, nil)

	pr := models.Principal("")

	response := handlers.HandlerGroups(suite.mockDatabase, noodle_api.NewGetNoodleGroupsParams(), &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`[{"DN":"cn=admins,ou=groups,dc=example,dc=nz","Id":1,"Name":"Admins"},{"DN":"cn=users,ou=groups,dc=example,dc=nz","Id":2,"Name":"Users"}]`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *GroupHandlersTestSuite) TestHandlerGroupsGetOne() {
	suite.mockGroupTable.EXPECT().GetID(int64(2)).Return(models.Group{
		DN:   "cn=admins,ou=groups,dc=example,dc=nz",
		ID:   1,
		Name: "Admins",
	},
		nil)

	pr := models.Principal("")
	params := noodle_api.NewGetNoodleGroupsParams()
	var Groupid = int64(2)
	params.Groupid = &Groupid

	response := handlers.HandlerGroups(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`[{"DN":"cn=admins,ou=groups,dc=example,dc=nz","Id":1,"Name":"Admins"}]`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *GroupHandlersTestSuite) TestHandlerGroupsDBError() {
	suite.mockGroupTable.EXPECT().GetID(int64(2)).Return(models.Group{},
		errors.New("DB Failed"))

	pr := models.Principal("")
	params := noodle_api.NewGetNoodleGroupsParams()
	var Groupid = int64(2)
	params.Groupid = &Groupid

	response := handlers.HandlerGroups(suite.mockDatabase, params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriterTest(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()

	mockWriter.EXPECT().Write([]byte(`{"message":"DB Failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func TestGroupHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(GroupHandlersTestSuite))
}
