package handlers_test

import (
	"errors"
	"testing"

	"github.com/mwinters-stuff/noodle/handlers"
	"github.com/mwinters-stuff/noodle/noodle/database/mocks"
	ldap_mocks "github.com/mwinters-stuff/noodle/noodle/ldap_handler/mocks"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type LdapHandlersUserGroupsSyncTestSuite struct {
	suite.Suite
	mockLdap            *ldap_mocks.LdapHandler
	mockDatabase        *mocks.Database
	mockTables          *mocks.Tables
	mockGroupTable      *mocks.GroupTable
	mockUserTable       *mocks.UserTable
	mockUserGroupsTable *mocks.UserGroupsTable
}

func (suite *LdapHandlersUserGroupsSyncTestSuite) SetupSuite() {

}

func (suite *LdapHandlersUserGroupsSyncTestSuite) SetupTest() {
	suite.mockLdap = ldap_mocks.NewLdapHandler(suite.T())

	suite.mockDatabase = mocks.NewDatabase(suite.T())
	suite.mockTables = mocks.NewTables(suite.T())
	suite.mockGroupTable = mocks.NewGroupTable(suite.T())
	suite.mockUserTable = mocks.NewUserTable(suite.T())
	suite.mockUserGroupsTable = mocks.NewUserGroupsTable(suite.T())

}

func (suite *LdapHandlersUserGroupsSyncTestSuite) TearDownTest() {

}
func (suite *LdapHandlersUserGroupsSyncTestSuite) TearSuite() {

}

func (suite *LdapHandlersUserGroupsSyncTestSuite) TestHandlerSyncUserGroups() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(6)
	suite.mockTables.EXPECT().GroupTable().Return(suite.mockGroupTable).Times(2)
	suite.mockTables.EXPECT().UserTable().Return(suite.mockUserTable).Times(1)
	suite.mockTables.EXPECT().UserGroupsTable().Return(suite.mockUserGroupsTable).Times(3)

	suite.mockUserTable.EXPECT().GetAll().Once().Return([]*models.User{
		{
			ID:          2,
			DN:          "CN=jack,DC=example,DC=nz",
			Username:    "jack",
			DisplayName: "Jack M",
			Surname:     "M",
			GivenName:   "Jack",
			UIDNumber:   1002,
		},
	}, nil)

	suite.mockLdap.EXPECT().GetUserGroups(models.User{

		ID:          2,
		DN:          "CN=jack,DC=example,DC=nz",
		Username:    "jack",
		DisplayName: "Jack M",
		Surname:     "M",
		GivenName:   "Jack",
		UIDNumber:   1002,
	}).Return([]models.UserGroup{
		{
			GroupDN:   "cn=admins,ou=groups,dc=example,dc=nz",
			GroupName: "Admins",
			UserDN:    "CN=jack,DC=example,DC=nz",
			UserName:  "jack",
		},
		{
			GroupDN:   "cn=users,ou=groups,dc=example,dc=nz",
			GroupName: "Users",
			UserDN:    "CN=jack,DC=example,DC=nz",
			UserName:  "jack",
		},
	}, nil)

	suite.mockUserGroupsTable.EXPECT().GetUser(int64(2)).Return([]*models.UserGroup{
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
	}, nil).Times(1)

	suite.mockGroupTable.EXPECT().GetDN("cn=admins,ou=groups,dc=example,dc=nz").Once().Return(models.Group{
		ID:   1,
		DN:   "cn=admins,ou=groups,dc=example,dc=nz",
		Name: "Admins",
	}, nil)

	suite.mockGroupTable.EXPECT().GetDN("cn=users,ou=groups,dc=example,dc=nz").Once().Return(models.Group{
		ID:   2,
		DN:   "cn=users,ou=groups,dc=example,dc=nz",
		Name: "Users",
	}, nil)

	suite.mockUserGroupsTable.EXPECT().Insert(&models.UserGroup{
		GroupDN:   "cn=users,ou=groups,dc=example,dc=nz",
		GroupName: "Users",
		UserDN:    "CN=jack,DC=example,DC=nz",
		UserName:  "jack",
		UserID:    2,
		GroupID:   2,
	}).Return(nil)

	suite.mockUserGroupsTable.EXPECT().Delete(models.UserGroup{
		ID:        2,
		GroupID:   3,
		UserID:    2,
		GroupDN:   "cn=people,ou=groups,dc=example,dc=nz",
		GroupName: "People",
		UserDN:    "CN=jack,DC=example,DC=nz",
		UserName:  "jack",
	}).Once().Return(nil)

	response := handlers.SyncLDAPUserGroups(suite.mockDatabase, suite.mockLdap)
	require.Nil(suite.T(), response)

}

func (suite *LdapHandlersUserSyncTestSuite) TestHandlerSyncUserGroups_DBGetAllError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(1)
	suite.mockTables.EXPECT().UserTable().Return(suite.mockUserTable).Times(1)

	suite.mockUserTable.EXPECT().GetAll().Return([]*models.User{}, errors.New("failed")).Once()

	response := handlers.SyncLDAPUserGroups(suite.mockDatabase, suite.mockLdap)
	require.NotNil(suite.T(), response)
	require.Equal(suite.T(), noodle_api.NewGetNoodleLdapReloadConflict().WithPayload(&models.Error{Message: "failed"}), response)
}

func (suite *LdapHandlersUserSyncTestSuite) TestHandlerSyncUserGroups_LDAPGetUserGroupsError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(1)
	suite.mockTables.EXPECT().UserTable().Return(suite.mockUserTable).Times(1)

	suite.mockUserTable.EXPECT().GetAll().Once().Return([]*models.User{
		{
			ID:          2,
			DN:          "CN=jack,DC=example,DC=nz",
			Username:    "jack",
			DisplayName: "Jack M",
			Surname:     "M",
			GivenName:   "Jack",
			UIDNumber:   1002,
		},
	}, nil)

	suite.mockLdap.EXPECT().GetUserGroups(models.User{
		ID:          2,
		DN:          "CN=jack,DC=example,DC=nz",
		Username:    "jack",
		DisplayName: "Jack M",
		Surname:     "M",
		GivenName:   "Jack",
		UIDNumber:   1002,
	}).Return([]models.UserGroup{}, errors.New("failed")).Once()

	response := handlers.SyncLDAPUserGroups(suite.mockDatabase, suite.mockLdap)
	require.NotNil(suite.T(), response)
	require.Equal(suite.T(), noodle_api.NewGetNoodleLdapReloadConflict().WithPayload(&models.Error{Message: "failed"}), response)
}

func (suite *LdapHandlersUserSyncTestSuite) TestHandlerSyncUserGroups_DBGetUserGroupsError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(2)
	suite.mockTables.EXPECT().UserTable().Return(suite.mockUserTable).Times(1)
	suite.mockTables.EXPECT().UserGroupsTable().Return(suite.mockUserGroupsTable).Times(1)

	suite.mockUserTable.EXPECT().GetAll().Once().Return([]*models.User{
		{
			ID:          2,
			DN:          "CN=jack,DC=example,DC=nz",
			Username:    "jack",
			DisplayName: "Jack M",
			Surname:     "M",
			GivenName:   "Jack",
			UIDNumber:   1002,
		},
	}, nil)

	suite.mockLdap.EXPECT().GetUserGroups(models.User{
		ID:          2,
		DN:          "CN=jack,DC=example,DC=nz",
		Username:    "jack",
		DisplayName: "Jack M",
		Surname:     "M",
		GivenName:   "Jack",
		UIDNumber:   1002,
	}).Return([]models.UserGroup{
		{
			GroupDN:   "cn=admins,ou=groups,dc=example,dc=nz",
			GroupName: "Admins",
			UserDN:    "CN=jack,DC=example,DC=nz",
			UserName:  "jack",
		},
		{
			GroupDN:   "cn=users,ou=groups,dc=example,dc=nz",
			GroupName: "Users",
			UserDN:    "CN=jack,DC=example,DC=nz",
			UserName:  "jack",
		},
	}, nil)

	suite.mockUserGroupsTable.EXPECT().GetUser(int64(2)).Return([]*models.UserGroup{}, errors.New("failed")).Once()

	response := handlers.SyncLDAPUserGroups(suite.mockDatabase, suite.mockLdap)
	require.NotNil(suite.T(), response)
	require.Equal(suite.T(), noodle_api.NewGetNoodleLdapReloadConflict().WithPayload(&models.Error{Message: "failed"}), response)
}

func (suite *LdapHandlersUserSyncTestSuite) TestHandlerSyncUserGroups_DBGetGroupDNError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(3)
	suite.mockTables.EXPECT().GroupTable().Return(suite.mockGroupTable).Times(1)
	suite.mockTables.EXPECT().UserTable().Return(suite.mockUserTable).Times(1)
	suite.mockTables.EXPECT().UserGroupsTable().Return(suite.mockUserGroupsTable).Times(1)

	suite.mockUserTable.EXPECT().GetAll().Once().Return([]*models.User{
		{
			ID:          2,
			DN:          "CN=jack,DC=example,DC=nz",
			Username:    "jack",
			DisplayName: "Jack M",
			Surname:     "M",
			GivenName:   "Jack",
			UIDNumber:   1002,
		},
	}, nil)

	suite.mockLdap.EXPECT().GetUserGroups(models.User{
		ID:          2,
		DN:          "CN=jack,DC=example,DC=nz",
		Username:    "jack",
		DisplayName: "Jack M",
		Surname:     "M",
		GivenName:   "Jack",
		UIDNumber:   1002,
	}).Return([]models.UserGroup{
		{
			GroupDN:   "cn=admins,ou=groups,dc=example,dc=nz",
			GroupName: "Admins",
			UserDN:    "CN=jack,DC=example,DC=nz",
			UserName:  "jack",
		},
		{
			GroupDN:   "cn=users,ou=groups,dc=example,dc=nz",
			GroupName: "Users",
			UserDN:    "CN=jack,DC=example,DC=nz",
			UserName:  "jack",
		},
	}, nil)

	suite.mockUserGroupsTable.EXPECT().GetUser(int64(2)).Return([]*models.UserGroup{
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
	}, nil).Times(1)

	suite.mockGroupTable.EXPECT().GetDN("cn=admins,ou=groups,dc=example,dc=nz").Once().Return(models.Group{}, errors.New("failed")).Once()

	response := handlers.SyncLDAPUserGroups(suite.mockDatabase, suite.mockLdap)
	require.NotNil(suite.T(), response)
	require.Equal(suite.T(), noodle_api.NewGetNoodleLdapReloadConflict().WithPayload(&models.Error{Message: "failed"}), response)
}

func (suite *LdapHandlersUserSyncTestSuite) TestHandlerSyncUserGroups_DBInsertUserGroupError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(5)
	suite.mockTables.EXPECT().GroupTable().Return(suite.mockGroupTable).Times(2)
	suite.mockTables.EXPECT().UserTable().Return(suite.mockUserTable).Times(1)
	suite.mockTables.EXPECT().UserGroupsTable().Return(suite.mockUserGroupsTable).Times(2)

	suite.mockUserTable.EXPECT().GetAll().Once().Return([]*models.User{
		{
			ID:          2,
			DN:          "CN=jack,DC=example,DC=nz",
			Username:    "jack",
			DisplayName: "Jack M",
			Surname:     "M",
			GivenName:   "Jack",
			UIDNumber:   1002,
		},
	}, nil)

	suite.mockLdap.EXPECT().GetUserGroups(models.User{
		ID:          2,
		DN:          "CN=jack,DC=example,DC=nz",
		Username:    "jack",
		DisplayName: "Jack M",
		Surname:     "M",
		GivenName:   "Jack",
		UIDNumber:   1002,
	}).Return([]models.UserGroup{
		{
			GroupDN:   "cn=admins,ou=groups,dc=example,dc=nz",
			GroupName: "Admins",
			UserDN:    "CN=jack,DC=example,DC=nz",
			UserName:  "jack",
		},
		{
			GroupDN:   "cn=users,ou=groups,dc=example,dc=nz",
			GroupName: "Users",
			UserDN:    "CN=jack,DC=example,DC=nz",
			UserName:  "jack",
		},
	}, nil)

	suite.mockUserGroupsTable.EXPECT().GetUser(int64(2)).Return([]*models.UserGroup{
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
	}, nil).Times(1)

	suite.mockGroupTable.EXPECT().GetDN("cn=admins,ou=groups,dc=example,dc=nz").Once().Return(models.Group{
		ID:   1,
		DN:   "cn=admins,ou=groups,dc=example,dc=nz",
		Name: "Admins",
	}, nil)

	suite.mockGroupTable.EXPECT().GetDN("cn=users,ou=groups,dc=example,dc=nz").Once().Return(models.Group{
		ID:   2,
		DN:   "cn=users,ou=groups,dc=example,dc=nz",
		Name: "Users",
	}, nil)

	suite.mockUserGroupsTable.EXPECT().Insert(&models.UserGroup{
		GroupDN:   "cn=users,ou=groups,dc=example,dc=nz",
		GroupName: "Users",
		UserDN:    "CN=jack,DC=example,DC=nz",
		UserName:  "jack",
		UserID:    2,
		GroupID:   2,
	}).Return(errors.New("failed"))

	response := handlers.SyncLDAPUserGroups(suite.mockDatabase, suite.mockLdap)
	require.NotNil(suite.T(), response)
	require.Equal(suite.T(), noodle_api.NewGetNoodleLdapReloadConflict().WithPayload(&models.Error{Message: "failed"}), response)
}

func (suite *LdapHandlersUserSyncTestSuite) TestHandlerSyncUserGroups_DBDeleteUserGroupError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(6)
	suite.mockTables.EXPECT().GroupTable().Return(suite.mockGroupTable).Times(2)
	suite.mockTables.EXPECT().UserTable().Return(suite.mockUserTable).Times(1)
	suite.mockTables.EXPECT().UserGroupsTable().Return(suite.mockUserGroupsTable).Times(3)

	suite.mockUserTable.EXPECT().GetAll().Once().Return([]*models.User{
		{
			ID:          2,
			DN:          "CN=jack,DC=example,DC=nz",
			Username:    "jack",
			DisplayName: "Jack M",
			Surname:     "M",
			GivenName:   "Jack",
			UIDNumber:   1002,
		},
	}, nil)

	suite.mockLdap.EXPECT().GetUserGroups(models.User{
		ID:          2,
		DN:          "CN=jack,DC=example,DC=nz",
		Username:    "jack",
		DisplayName: "Jack M",
		Surname:     "M",
		GivenName:   "Jack",
		UIDNumber:   1002,
	}).Return([]models.UserGroup{
		{
			GroupDN:   "cn=admins,ou=groups,dc=example,dc=nz",
			GroupName: "Admins",
			UserDN:    "CN=jack,DC=example,DC=nz",
			UserName:  "jack",
		},
		{
			GroupDN:   "cn=users,ou=groups,dc=example,dc=nz",
			GroupName: "Users",
			UserDN:    "CN=jack,DC=example,DC=nz",
			UserName:  "jack",
		},
	}, nil)

	suite.mockUserGroupsTable.EXPECT().GetUser(int64(2)).Return([]*models.UserGroup{
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
	}, nil).Times(1)

	suite.mockGroupTable.EXPECT().GetDN("cn=admins,ou=groups,dc=example,dc=nz").Once().Return(models.Group{
		ID:   1,
		DN:   "cn=admins,ou=groups,dc=example,dc=nz",
		Name: "Admins",
	}, nil)

	suite.mockGroupTable.EXPECT().GetDN("cn=users,ou=groups,dc=example,dc=nz").Once().Return(models.Group{
		ID:   2,
		DN:   "cn=users,ou=groups,dc=example,dc=nz",
		Name: "Users",
	}, nil)

	suite.mockUserGroupsTable.EXPECT().Insert(&models.UserGroup{
		GroupDN:   "cn=users,ou=groups,dc=example,dc=nz",
		GroupName: "Users",
		UserDN:    "CN=jack,DC=example,DC=nz",
		UserName:  "jack",
		UserID:    2,
		GroupID:   2,
	}).Return(nil)

	suite.mockUserGroupsTable.EXPECT().Delete(models.UserGroup{
		ID:        2,
		GroupID:   3,
		UserID:    2,
		GroupDN:   "cn=people,ou=groups,dc=example,dc=nz",
		GroupName: "People",
		UserDN:    "CN=jack,DC=example,DC=nz",
		UserName:  "jack",
	}).Return(errors.New("failed"))

	response := handlers.SyncLDAPUserGroups(suite.mockDatabase, suite.mockLdap)
	require.NotNil(suite.T(), response)
	require.Equal(suite.T(), noodle_api.NewGetNoodleLdapReloadConflict().WithPayload(&models.Error{Message: "failed"}), response)
}

func TestLdapHandlersUserGroupsSyncTestSuite(t *testing.T) {
	suite.Run(t, new(LdapHandlersUserGroupsSyncTestSuite))
}
