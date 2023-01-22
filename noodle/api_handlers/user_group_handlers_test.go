package api_handlers_test

import (
	"errors"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/mwinters-stuff/noodle/noodle/api_handlers"
	handler_mocks "github.com/mwinters-stuff/noodle/noodle/api_handlers/mocks"
	"github.com/mwinters-stuff/noodle/noodle/database/mocks"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/ldap_handler"
	"github.com/rs/zerolog/log"
)

type UserGroupHandlersTestSuite struct {
	suite.Suite
	mockDatabase       *mocks.Database
	mockTables         *mocks.Tables
	mockUserGroupTable *mocks.UserGroupsTable
	api                *operations.NoodleAPI
}

func (suite *UserGroupHandlersTestSuite) SetupSuite() {
	database.Logger = log.Output(nil)
	ldap_handler.Logger = log.Output(nil)
}

func (suite *UserGroupHandlersTestSuite) SetupTest() {
	suite.mockDatabase = mocks.NewDatabase(suite.T())
	suite.mockTables = mocks.NewTables(suite.T())
	suite.mockUserGroupTable = mocks.NewUserGroupsTable(suite.T())

	suite.api = &operations.NoodleAPI{}
	api_handlers.RegisterUserGroupApiHandlers(suite.api, suite.mockDatabase)

	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleUserGroupsHandler)
}

func (suite *UserGroupHandlersTestSuite) TearDownTest() {

}
func (suite *UserGroupHandlersTestSuite) TearSuite() {

}

func (suite *UserGroupHandlersTestSuite) TestHandlerGetGroupUsers() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(1)
	suite.mockTables.EXPECT().UserGroupsTable().Return(suite.mockUserGroupTable).Times(1)

	suite.mockUserGroupTable.EXPECT().GetGroup(int64(1)).Return([]*models.UserGroup{
		{
			ID:        1,
			GroupID:   1,
			UserID:    1,
			GroupDN:   "cn=admins,ou=groups,dc=example,dc=nz",
			GroupName: "Admins",
			UserDN:    "CN=bob,DC=example,DC=nz",
			UserName:  "bobe",
		},
		{
			ID:        2,
			GroupID:   1,
			UserID:    2,
			GroupDN:   "cn=admins,ou=groups,dc=example,dc=nz",
			GroupName: "Admins",
			UserDN:    "CN=jack,DC=example,DC=nz",
			UserName:  "jack",
		},
	},
		nil)

	pr := models.Principal("")
	params := noodle_api.NewGetNoodleUserGroupsParams()
	var Groupid = int64(1)
	params.Groupid = &Groupid

	response := suite.api.NoodleAPIGetNoodleUserGroupsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`[{"GroupDN":"cn=admins,ou=groups,dc=example,dc=nz","GroupId":1,"GroupName":"Admins","Id":1,"UserDN":"CN=bob,DC=example,DC=nz","UserId":1,"UserName":"bobe"},{"GroupDN":"cn=admins,ou=groups,dc=example,dc=nz","GroupId":1,"GroupName":"Admins","Id":2,"UserDN":"CN=jack,DC=example,DC=nz","UserId":2,"UserName":"jack"}]`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *UserGroupHandlersTestSuite) TestHandlerGetUserGroups() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(1)
	suite.mockTables.EXPECT().UserGroupsTable().Return(suite.mockUserGroupTable).Times(1)

	suite.mockUserGroupTable.EXPECT().GetUser(int64(1)).Return([]*models.UserGroup{
		{
			ID:        1,
			GroupID:   1,
			UserID:    1,
			GroupDN:   "cn=admins,ou=groups,dc=example,dc=nz",
			GroupName: "Admins",
			UserDN:    "CN=bob,DC=example,DC=nz",
			UserName:  "bobe",
		},
		{
			ID:        2,
			GroupID:   2,
			UserID:    1,
			GroupDN:   "cn=users,ou=groups,dc=example,dc=nz",
			GroupName: "Users",
			UserDN:    "CN=bob,DC=example,DC=nz",
			UserName:  "bobe",
		},
	},
		nil)

	pr := models.Principal("")
	params := noodle_api.NewGetNoodleUserGroupsParams()
	var Userid = int64(1)
	params.Userid = &Userid

	response := suite.api.NoodleAPIGetNoodleUserGroupsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(200).Once()

	mockWriter.EXPECT().Write([]byte(`[{"GroupDN":"cn=admins,ou=groups,dc=example,dc=nz","GroupId":1,"GroupName":"Admins","Id":1,"UserDN":"CN=bob,DC=example,DC=nz","UserId":1,"UserName":"bobe"},{"GroupDN":"cn=users,ou=groups,dc=example,dc=nz","GroupId":2,"GroupName":"Users","Id":2,"UserDN":"CN=bob,DC=example,DC=nz","UserId":1,"UserName":"bobe"}]`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *UserGroupHandlersTestSuite) TestHandlerGroupsDBError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(1)
	suite.mockTables.EXPECT().UserGroupsTable().Return(suite.mockUserGroupTable).Times(1)

	suite.mockUserGroupTable.EXPECT().GetGroup(int64(2)).Return([]*models.UserGroup{},
		errors.New("DB Failed"))

	pr := models.Principal("")
	params := noodle_api.NewGetNoodleUserGroupsParams()
	var Groupid = int64(2)
	params.Groupid = &Groupid

	response := suite.api.NoodleAPIGetNoodleUserGroupsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()

	mockWriter.EXPECT().Write([]byte(`{"message":"DB Failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *UserGroupHandlersTestSuite) TestHandlerTwoParametersError() {

	pr := models.Principal("")
	params := noodle_api.NewGetNoodleUserGroupsParams()

	params.Groupid = nil
	params.Userid = nil

	response := suite.api.NoodleAPIGetNoodleUserGroupsHandler.Handle(params, &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())
	mockWriter.EXPECT().WriteHeader(409).Once()

	mockWriter.EXPECT().Write([]byte(`{"message":"no groupid or userid parameter"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func TestUserGroupHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(UserGroupHandlersTestSuite))
}
