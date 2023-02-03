package api_handlers_test

import (
	"errors"
	"testing"

	"github.com/mwinters-stuff/noodle/noodle/api_handlers"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/database/mocks"
	"github.com/mwinters-stuff/noodle/noodle/ldap_handler"
	ldap_mocks "github.com/mwinters-stuff/noodle/noodle/ldap_handler/mocks"
	"github.com/mwinters-stuff/noodle/server/models"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type LdapHandlersGroupSyncTestSuite struct {
	suite.Suite
	mockLdap            *ldap_mocks.LdapHandler
	mockDatabase        *mocks.Database
	mockTables          *mocks.Tables
	mockGroupTable      *mocks.GroupTable
	mockUserTable       *mocks.UserTable
	mockUserGroupsTable *mocks.UserGroupsTable
}

func (suite *LdapHandlersGroupSyncTestSuite) SetupSuite() {
	database.Logger = log.Output(nil)
	ldap_handler.Logger = log.Output(nil)
	api_handlers.Logger = log.Output(nil)
}

func (suite *LdapHandlersGroupSyncTestSuite) SetupTest() {
	suite.mockLdap = ldap_mocks.NewLdapHandler(suite.T())

	suite.mockDatabase = mocks.NewDatabase(suite.T())
	suite.mockTables = mocks.NewTables(suite.T())
	suite.mockGroupTable = mocks.NewGroupTable(suite.T())
	suite.mockUserTable = mocks.NewUserTable(suite.T())
	suite.mockUserGroupsTable = mocks.NewUserGroupsTable(suite.T())

}

func (suite *LdapHandlersGroupSyncTestSuite) TearDownTest() {

}
func (suite *LdapHandlersGroupSyncTestSuite) TearDownSuite() {

}

func (suite *LdapHandlersGroupSyncTestSuite) TestHandlerSyncGroups() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(6)
	suite.mockTables.EXPECT().GroupTable().Return(suite.mockGroupTable).Times(6)

	suite.mockLdap.EXPECT().GetGroups().Return([]models.Group{
		{
			DN:   "cn=admins,ou=groups,dc=example,dc=nz",
			Name: "Admins",
		},
		{
			DN:   "cn=users,ou=groups,dc=example,dc=nz",
			Name: "Users",
		},
	}, nil)

	suite.mockGroupTable.EXPECT().GetAll().Return([]*models.Group{
		{
			ID:   1,
			DN:   "cn=admins,ou=groups,dc=example,dc=nz",
			Name: "Admins",
		},
		{
			ID:   2,
			DN:   "cn=people,ou=groups,dc=example,dc=nz",
			Name: "People",
		},
	}, nil).Times(1)

	suite.mockGroupTable.EXPECT().ExistsDN("cn=admins,ou=groups,dc=example,dc=nz").Once().Return(true, nil)
	suite.mockGroupTable.EXPECT().Update(models.Group{
		ID:   1,
		DN:   "cn=admins,ou=groups,dc=example,dc=nz",
		Name: "Admins",
	}).Once().Return(nil)

	suite.mockGroupTable.EXPECT().ExistsDN("cn=users,ou=groups,dc=example,dc=nz").Once().Return(false, nil)

	suite.mockGroupTable.EXPECT().Delete(models.Group{
		ID:   2,
		DN:   "cn=people,ou=groups,dc=example,dc=nz",
		Name: "People",
	}).Once().Return(nil)

	suite.mockGroupTable.EXPECT().Insert(&models.Group{
		DN:   "cn=users,ou=groups,dc=example,dc=nz",
		Name: "Users",
	}).Return(nil)

	response := api_handlers.SyncLDAPGroups(suite.mockDatabase, suite.mockLdap)
	require.Nil(suite.T(), response)

}

func (suite *LdapHandlersGroupSyncTestSuite) TestHandlerSyncGroups_LdapGetGroupsError() {

	suite.mockLdap.EXPECT().GetGroups().Return([]models.Group{}, errors.New("failed"))

	response := api_handlers.SyncLDAPGroups(suite.mockDatabase, suite.mockLdap)
	require.NotNil(suite.T(), response)
	require.ErrorContains(suite.T(), response, "failed")
}

func (suite *LdapHandlersGroupSyncTestSuite) TestHandlerSyncGroups_DBGetAllError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(1)
	suite.mockTables.EXPECT().GroupTable().Return(suite.mockGroupTable).Times(1)

	suite.mockLdap.EXPECT().GetGroups().Return([]models.Group{
		{
			DN:   "cn=admins,ou=groups,dc=example,dc=nz",
			Name: "Admins",
		},
	}, nil)

	suite.mockGroupTable.EXPECT().GetAll().Return([]*models.Group{}, errors.New("failed"))

	response := api_handlers.SyncLDAPGroups(suite.mockDatabase, suite.mockLdap)
	require.NotNil(suite.T(), response)
	require.ErrorContains(suite.T(), response, "failed")
}

func (suite *LdapHandlersGroupSyncTestSuite) TestHandlerSyncGroups_DBExistsDNError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(2)
	suite.mockTables.EXPECT().GroupTable().Return(suite.mockGroupTable).Times(2)

	suite.mockLdap.EXPECT().GetGroups().Return([]models.Group{
		{
			DN:   "cn=admins,ou=groups,dc=example,dc=nz",
			Name: "Admins",
		},
	}, nil)

	suite.mockGroupTable.EXPECT().GetAll().Return([]*models.Group{
		{
			ID:   1,
			DN:   "cn=admins,ou=groups,dc=example,dc=nz",
			Name: "Admins",
		},
	}, nil)

	suite.mockGroupTable.EXPECT().ExistsDN("cn=admins,ou=groups,dc=example,dc=nz").Once().Return(false, errors.New("failed"))

	response := api_handlers.SyncLDAPGroups(suite.mockDatabase, suite.mockLdap)
	require.NotNil(suite.T(), response)
	require.ErrorContains(suite.T(), response, "failed")
}

func (suite *LdapHandlersGroupSyncTestSuite) TestHandlerSyncGroups_DBUpdateError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(3)
	suite.mockTables.EXPECT().GroupTable().Return(suite.mockGroupTable).Times(3)

	suite.mockLdap.EXPECT().GetGroups().Return([]models.Group{
		{
			DN:   "cn=admins,ou=groups,dc=example,dc=nz",
			Name: "Admins",
		},
	}, nil)

	suite.mockGroupTable.EXPECT().GetAll().Return([]*models.Group{
		{
			ID:   1,
			DN:   "cn=admins,ou=groups,dc=example,dc=nz",
			Name: "Admins",
		},
	}, nil)

	suite.mockGroupTable.EXPECT().ExistsDN("cn=admins,ou=groups,dc=example,dc=nz").Once().Return(true, nil)

	suite.mockGroupTable.EXPECT().Update(models.Group{
		ID:   1,
		DN:   "cn=admins,ou=groups,dc=example,dc=nz",
		Name: "Admins",
	}).Once().Return(errors.New("failed"))

	response := api_handlers.SyncLDAPGroups(suite.mockDatabase, suite.mockLdap)
	require.NotNil(suite.T(), response)
	require.ErrorContains(suite.T(), response, "failed")
}

func (suite *LdapHandlersGroupSyncTestSuite) TestHandlerSyncGroups_DBDeleteError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(2)
	suite.mockTables.EXPECT().GroupTable().Return(suite.mockGroupTable).Times(2)

	suite.mockLdap.EXPECT().GetGroups().Return([]models.Group{}, nil)

	suite.mockGroupTable.EXPECT().GetAll().Return([]*models.Group{
		{
			ID:   1,
			DN:   "cn=admins,ou=groups,dc=example,dc=nz",
			Name: "Admins",
		},
	}, nil)

	suite.mockGroupTable.EXPECT().Delete(models.Group{
		ID:   1,
		DN:   "cn=admins,ou=groups,dc=example,dc=nz",
		Name: "Admins",
	}).Once().Return(errors.New("failed"))

	response := api_handlers.SyncLDAPGroups(suite.mockDatabase, suite.mockLdap)
	require.NotNil(suite.T(), response)
	require.ErrorContains(suite.T(), response, "failed")
}

func (suite *LdapHandlersGroupSyncTestSuite) TestHandlerSyncGroups_DBInsertError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(3)
	suite.mockTables.EXPECT().GroupTable().Return(suite.mockGroupTable).Times(3)

	suite.mockLdap.EXPECT().GetGroups().Return([]models.Group{
		{
			DN:   "cn=admins,ou=groups,dc=example,dc=nz",
			Name: "Admins",
		},
	}, nil)

	suite.mockGroupTable.EXPECT().GetAll().Return([]*models.Group{}, nil)

	suite.mockGroupTable.EXPECT().ExistsDN("cn=admins,ou=groups,dc=example,dc=nz").Once().Return(false, nil)

	suite.mockGroupTable.EXPECT().Insert(&models.Group{
		DN:   "cn=admins,ou=groups,dc=example,dc=nz",
		Name: "Admins",
	}).Once().Return(errors.New("failed"))

	response := api_handlers.SyncLDAPGroups(suite.mockDatabase, suite.mockLdap)
	require.NotNil(suite.T(), response)
	require.ErrorContains(suite.T(), response, "failed")
}

func TestLdapHandlersGroupSyncTestSuite(t *testing.T) {
	suite.Run(t, new(LdapHandlersGroupSyncTestSuite))
}
