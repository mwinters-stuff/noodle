package api_handlers_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/rs/zerolog/log"

	"github.com/mwinters-stuff/noodle/noodle/api_handlers"
	handler_mocks "github.com/mwinters-stuff/noodle/noodle/api_handlers/mocks"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/database/mocks"
	"github.com/mwinters-stuff/noodle/noodle/ldap_handler"
	ldap_mocks "github.com/mwinters-stuff/noodle/noodle/ldap_handler/mocks"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type LdapHandlersOtherTestSuite struct {
	suite.Suite
	mockLdap            *ldap_mocks.LdapHandler
	mockDatabase        *mocks.Database
	mockTables          *mocks.Tables
	mockGroupTable      *mocks.GroupTable
	mockUserTable       *mocks.UserTable
	mockUserGroupsTable *mocks.UserGroupsTable
	api                 *operations.NoodleAPI
}

func (suite *LdapHandlersOtherTestSuite) SetupSuite() {
	database.Logger = log.Output(nil)
	ldap_handler.Logger = log.Output(nil)
	api_handlers.Logger = log.Output(nil)
}

func (suite *LdapHandlersOtherTestSuite) SetupTest() {
	suite.mockLdap = ldap_mocks.NewLdapHandler(suite.T())

	suite.mockDatabase = mocks.NewDatabase(suite.T())
	suite.mockTables = mocks.NewTables(suite.T())
	suite.mockGroupTable = mocks.NewGroupTable(suite.T())
	suite.mockUserTable = mocks.NewUserTable(suite.T())
	suite.mockUserGroupsTable = mocks.NewUserGroupsTable(suite.T())

	suite.api = &operations.NoodleAPI{}
	api_handlers.RegisterLdapApiHandlers(suite.api, suite.mockDatabase, suite.mockLdap)

	require.NotNil(suite.T(), suite.api.NoodleAPIGetNoodleLdapReloadHandler)

}

func (suite *LdapHandlersOtherTestSuite) TearDownTest() {

}
func (suite *LdapHandlersOtherTestSuite) TearDownSuite() {

}

func (suite *LdapHandlersOtherTestSuite) TestHandlerNoUsersOrGroupsLdapOrDatabase() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(3)

	suite.mockLdap.EXPECT().GetUsers().Return([]models.User{}, nil)

	suite.mockTables.EXPECT().UserTable().Return(suite.mockUserTable).Times(1)
	suite.mockUserTable.EXPECT().GetAll().Return([]*models.User{}, nil).Times(1)

	suite.mockLdap.EXPECT().GetGroups().Return([]models.Group{}, nil)

	suite.mockTables.EXPECT().GroupTable().Return(suite.mockGroupTable).Times(1)
	suite.mockGroupTable.EXPECT().GetAll().Return([]*models.Group{}, nil).Times(1)

	suite.mockTables.EXPECT().UserTable().Return(suite.mockUserTable).Times(1)
	suite.mockUserTable.EXPECT().GetAll().Return([]*models.User{}, nil).Times(1)

	pr := models.Principal("")
	response := suite.api.NoodleAPIGetNoodleLdapReloadHandler.Handle(noodle_api.NewGetNoodleLdapReloadParams(), &pr)

	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())

	mockWriter.EXPECT().Header().Once().Return(http.Header{})
	mockWriter.EXPECT().WriteHeader(200).Once()

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())
}

func (suite *LdapHandlersOtherTestSuite) TestHandlerUserSyncError() {
	suite.mockLdap.EXPECT().GetUsers().Return([]models.User{}, errors.New("failed")).Once()

	pr := models.Principal("")

	response := suite.api.NoodleAPIGetNoodleLdapReloadHandler.Handle(noodle_api.NewGetNoodleLdapReloadParams(), &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())

	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())

}

func (suite *LdapHandlersOtherTestSuite) TestHandlerGroupSyncError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(1)
	suite.mockTables.EXPECT().UserTable().Return(suite.mockUserTable).Times(1)

	suite.mockLdap.EXPECT().GetUsers().Return([]models.User{}, nil).Once()

	suite.mockUserTable.EXPECT().GetAll().Return([]*models.User{}, nil).Once()

	suite.mockLdap.EXPECT().GetGroups().Return([]models.Group{}, errors.New("failed")).Once()

	pr := models.Principal("")
	response := suite.api.NoodleAPIGetNoodleLdapReloadHandler.Handle(noodle_api.NewGetNoodleLdapReloadParams(), &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())

	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())

}

func (suite *LdapHandlersOtherTestSuite) TestHandlerUserGroupsSyncError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(3)
	suite.mockTables.EXPECT().UserTable().Return(suite.mockUserTable).Times(2)
	suite.mockTables.EXPECT().GroupTable().Return(suite.mockGroupTable).Times(1)

	suite.mockLdap.EXPECT().GetUsers().Return([]models.User{}, nil).Once()

	suite.mockUserTable.EXPECT().GetAll().Return([]*models.User{}, nil).Once()

	suite.mockLdap.EXPECT().GetGroups().Return([]models.Group{}, nil).Once()

	suite.mockGroupTable.EXPECT().GetAll().Return([]*models.Group{}, nil).Once()

	suite.mockUserTable.EXPECT().GetAll().Return([]*models.User{}, errors.New("failed")).Once()

	pr := models.Principal("")
	response := suite.api.NoodleAPIGetNoodleLdapReloadHandler.Handle(noodle_api.NewGetNoodleLdapReloadParams(), &pr)
	require.NotNil(suite.T(), response)

	mockWriter := handler_mocks.NewResponseWriter(suite.T())

	mockWriter.EXPECT().WriteHeader(409).Once()
	mockWriter.EXPECT().Write([]byte(`{"message":"failed"}`)).Once().Return(1, nil)

	response.WriteResponse(mockWriter, runtime.ByteStreamProducer())

}

func (suite *LdapHandlersOtherTestSuite) TestIndexUserGroups() {
	list := []*models.UserGroup{
		{
			ID:        1,
			GroupID:   1,
			UserID:    2,
			GroupDN:   "cn=admins,ou=groups,dc=example,dc=nz",
			GroupName: "Admins",
			UserDN:    "CN=jack,DC=example,DC=nz",
			UserName:  "jack",
		},
		{
			ID:        2,
			GroupID:   3,
			UserID:    2,
			GroupDN:   "cn=people,ou=groups,dc=example,dc=nz",
			GroupName: "People",
			UserDN:    "CN=jack,DC=example,DC=nz",
			UserName:  "jack",
		},
	}

	require.Equal(suite.T(), 1, api_handlers.IndexUserGroup(list, models.UserGroup{
		ID:        2,
		GroupID:   3,
		UserID:    2,
		GroupDN:   "cn=people,ou=groups,dc=example,dc=nz",
		GroupName: "People",
		UserDN:    "CN=jack,DC=example,DC=nz",
		UserName:  "jack",
	}))

	require.Equal(suite.T(), 0, api_handlers.IndexUserGroup(list, models.UserGroup{
		ID:        1,
		GroupID:   1,
		UserID:    2,
		GroupDN:   "cn=admins,ou=groups,dc=example,dc=nz",
		GroupName: "Admins",
		UserDN:    "CN=jack,DC=example,DC=nz",
		UserName:  "jack",
	}))
	require.Equal(suite.T(), -1, api_handlers.IndexUserGroup(list, models.UserGroup{
		ID:        1,
		GroupID:   9,
		UserID:    2,
		GroupDN:   "cn=admins,ou=groups,dc=example,dc=nz",
		GroupName: "Admins",
		UserDN:    "CN=jack,DC=example,DC=nz",
		UserName:  "jack",
	}))
}

func (suite *LdapHandlersOtherTestSuite) TestIndexGroups() {
	list := []*models.Group{
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
	}

	require.Equal(suite.T(), 1, api_handlers.IndexGroup(list, models.Group{
		ID:   2,
		DN:   "cn=users,ou=groups,dc=example,dc=nz",
		Name: "Users",
	}))

	require.Equal(suite.T(), 0, api_handlers.IndexGroup(list, models.Group{
		ID:   1,
		DN:   "cn=admins,ou=groups,dc=example,dc=nz",
		Name: "Admins",
	}))
	require.Equal(suite.T(), -1, api_handlers.IndexGroup(list, models.Group{
		ID:   3,
		DN:   "cn=people,ou=groups,dc=example,dc=nz",
		Name: "People",
	}))
}

func (suite *LdapHandlersOtherTestSuite) TestIndexUsers() {
	list := []*models.User{
		{
			ID:          2,
			DN:          "CN=jack,DC=example,DC=nz",
			Username:    "jack",
			DisplayName: "Jack M",
			Surname:     "M",
			GivenName:   "Jack",
			UIDNumber:   1002,
		},
		{
			ID:          1,
			DN:          "CN=bob,DC=example,DC=nz",
			Username:    "bobe",
			DisplayName: "bobextample",
			Surname:     "Extample",
			GivenName:   "Bob",
			UIDNumber:   1001,
		},
	}

	require.Equal(suite.T(), 1, api_handlers.IndexUser(list, models.User{
		ID:          1,
		DN:          "CN=bob,DC=example,DC=nz",
		Username:    "bobe",
		DisplayName: "bobextample",
		Surname:     "Extample",
		GivenName:   "Bob",
		UIDNumber:   1001,
	}))

	require.Equal(suite.T(), 0, api_handlers.IndexUser(list, models.User{
		ID:          2,
		DN:          "CN=jack,DC=example,DC=nz",
		Username:    "jack",
		DisplayName: "Jack M",
		Surname:     "M",
		GivenName:   "Jack",
		UIDNumber:   1002,
	}))
	require.Equal(suite.T(), -1, api_handlers.IndexUser(list, models.User{
		ID:          5,
		DN:          "CN=jill,DC=example,DC=nz",
		Username:    "jillie",
		DisplayName: "jilly",
		Surname:     "Frill",
		GivenName:   "Jill",
		UIDNumber:   1002,
	}))
}

func TestLdapHandlersOtherTestSuite(t *testing.T) {
	suite.Run(t, new(LdapHandlersOtherTestSuite))
}
