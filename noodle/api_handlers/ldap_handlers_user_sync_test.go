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
	"github.com/mwinters-stuff/noodle/server/restapi/operations/noodle_api"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type LdapHandlersUserSyncTestSuite struct {
	suite.Suite
	mockLdap            *ldap_mocks.LdapHandler
	mockDatabase        *mocks.Database
	mockTables          *mocks.Tables
	mockGroupTable      *mocks.GroupTable
	mockUserTable       *mocks.UserTable
	mockUserGroupsTable *mocks.UserGroupsTable
}

func (suite *LdapHandlersUserSyncTestSuite) SetupSuite() {
	database.Logger = log.Output(nil)
	ldap_handler.Logger = log.Output(nil)
	api_handlers.Logger = log.Output(nil)
}

func (suite *LdapHandlersUserSyncTestSuite) SetupTest() {
	suite.mockLdap = ldap_mocks.NewLdapHandler(suite.T())

	suite.mockDatabase = mocks.NewDatabase(suite.T())
	suite.mockTables = mocks.NewTables(suite.T())
	suite.mockGroupTable = mocks.NewGroupTable(suite.T())
	suite.mockUserTable = mocks.NewUserTable(suite.T())
	suite.mockUserGroupsTable = mocks.NewUserGroupsTable(suite.T())
}

func (suite *LdapHandlersUserSyncTestSuite) TearDownTest() {

}
func (suite *LdapHandlersUserSyncTestSuite) TearSuite() {

}

func (suite *LdapHandlersUserSyncTestSuite) TestHandlerSyncUsers() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(6)
	suite.mockTables.EXPECT().UserTable().Return(suite.mockUserTable).Times(6)

	suite.mockLdap.EXPECT().GetUsers().Return([]models.User{
		{
			DN:          "CN=bob,DC=example,DC=nz",
			Username:    "bobe",
			DisplayName: "bobextample",
			Surname:     "Extample",
			GivenName:   "Bob",
			UIDNumber:   1001,
		},
		{
			DN:          "CN=jill,DC=example,DC=nz",
			Username:    "jillie",
			DisplayName: "jilly",
			Surname:     "Frill",
			GivenName:   "Jill",
			UIDNumber:   1002,
		},
	}, nil)

	suite.mockUserTable.EXPECT().GetAll().Return([]*models.User{
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
	}, nil).Times(1)

	suite.mockUserTable.EXPECT().ExistsDN("CN=bob,DC=example,DC=nz").Once().Return(true, nil)
	suite.mockUserTable.EXPECT().Update(models.User{
		ID:          1,
		DN:          "CN=bob,DC=example,DC=nz",
		Username:    "bobe",
		DisplayName: "bobextample",
		Surname:     "Extample",
		GivenName:   "Bob",
		UIDNumber:   1001,
	}).Once().Return(nil)

	suite.mockUserTable.EXPECT().ExistsDN("CN=jill,DC=example,DC=nz").Once().Return(false, nil)

	suite.mockUserTable.EXPECT().Delete(models.User{
		ID:          2,
		DN:          "CN=jack,DC=example,DC=nz",
		Username:    "jack",
		DisplayName: "Jack M",
		Surname:     "M",
		GivenName:   "Jack",
		UIDNumber:   1002,
	}).Once().Return(nil)

	suite.mockUserTable.EXPECT().Insert(&models.User{
		DN:          "CN=jill,DC=example,DC=nz",
		Username:    "jillie",
		DisplayName: "jilly",
		Surname:     "Frill",
		GivenName:   "Jill",
		UIDNumber:   1002,
	}).Return(nil)

	response := api_handlers.SyncLDAPUsers(suite.mockDatabase, suite.mockLdap)
	require.Nil(suite.T(), response)

}

func (suite *LdapHandlersUserSyncTestSuite) TestHandlerSyncUsers_LdapGetUsersError() {

	suite.mockLdap.EXPECT().GetUsers().Return([]models.User{}, errors.New("failed"))

	response := api_handlers.SyncLDAPUsers(suite.mockDatabase, suite.mockLdap)
	require.NotNil(suite.T(), response)
	require.Equal(suite.T(), noodle_api.NewGetNoodleLdapReloadConflict().WithPayload(&models.Error{Message: "failed"}), response)
}

func (suite *LdapHandlersUserSyncTestSuite) TestHandlerSyncUsers_DBGetAllError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(1)
	suite.mockTables.EXPECT().UserTable().Return(suite.mockUserTable).Times(1)

	suite.mockLdap.EXPECT().GetUsers().Return([]models.User{
		{
			DN:          "CN=bob,DC=example,DC=nz",
			Username:    "bobe",
			DisplayName: "bobextample",
			Surname:     "Extample",
			GivenName:   "Bob",
			UIDNumber:   1001,
		},
	}, nil)

	suite.mockUserTable.EXPECT().GetAll().Return([]*models.User{}, errors.New("failed"))

	response := api_handlers.SyncLDAPUsers(suite.mockDatabase, suite.mockLdap)
	require.NotNil(suite.T(), response)
	require.Equal(suite.T(), noodle_api.NewGetNoodleLdapReloadConflict().WithPayload(&models.Error{Message: "failed"}), response)
}

func (suite *LdapHandlersUserSyncTestSuite) TestHandlerSyncUsers_DBExistsDNError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(2)
	suite.mockTables.EXPECT().UserTable().Return(suite.mockUserTable).Times(2)

	suite.mockLdap.EXPECT().GetUsers().Return([]models.User{
		{
			DN:          "CN=bob,DC=example,DC=nz",
			Username:    "bobe",
			DisplayName: "bobextample",
			Surname:     "Extample",
			GivenName:   "Bob",
			UIDNumber:   1001,
		},
	}, nil)

	suite.mockUserTable.EXPECT().GetAll().Return([]*models.User{
		{
			DN:          "CN=bob,DC=example,DC=nz",
			Username:    "bobe",
			DisplayName: "bobextample",
			Surname:     "Extample",
			GivenName:   "Bob",
			UIDNumber:   1001,
		},
	}, nil)

	suite.mockUserTable.EXPECT().ExistsDN("CN=bob,DC=example,DC=nz").Once().Return(false, errors.New("failed"))

	response := api_handlers.SyncLDAPUsers(suite.mockDatabase, suite.mockLdap)
	require.NotNil(suite.T(), response)
	require.Equal(suite.T(), noodle_api.NewGetNoodleLdapReloadConflict().WithPayload(&models.Error{Message: "failed"}), response)
}

func (suite *LdapHandlersUserSyncTestSuite) TestHandlerSyncUsers_DBUpdateError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(3)
	suite.mockTables.EXPECT().UserTable().Return(suite.mockUserTable).Times(3)

	suite.mockLdap.EXPECT().GetUsers().Return([]models.User{
		{
			DN:          "CN=bob,DC=example,DC=nz",
			Username:    "bobe",
			DisplayName: "bobextample",
			Surname:     "Extample",
			GivenName:   "Bob",
			UIDNumber:   1001,
		},
	}, nil)

	suite.mockUserTable.EXPECT().GetAll().Return([]*models.User{
		{
			DN:          "CN=bob,DC=example,DC=nz",
			ID:          1,
			Username:    "bobe",
			DisplayName: "bobextample",
			Surname:     "Extample",
			GivenName:   "Bob",
			UIDNumber:   1001,
		},
	}, nil)

	suite.mockUserTable.EXPECT().ExistsDN("CN=bob,DC=example,DC=nz").Once().Return(true, nil)

	suite.mockUserTable.EXPECT().Update(models.User{
		ID:          1,
		DN:          "CN=bob,DC=example,DC=nz",
		Username:    "bobe",
		DisplayName: "bobextample",
		Surname:     "Extample",
		GivenName:   "Bob",
		UIDNumber:   1001,
	}).Once().Return(errors.New("failed"))

	response := api_handlers.SyncLDAPUsers(suite.mockDatabase, suite.mockLdap)
	require.NotNil(suite.T(), response)
	require.Equal(suite.T(), noodle_api.NewGetNoodleLdapReloadConflict().WithPayload(&models.Error{Message: "failed"}), response)
}

func (suite *LdapHandlersUserSyncTestSuite) TestHandlerSyncUsers_DBDeleteError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(2)
	suite.mockTables.EXPECT().UserTable().Return(suite.mockUserTable).Times(2)

	suite.mockLdap.EXPECT().GetUsers().Return([]models.User{}, nil)

	suite.mockUserTable.EXPECT().GetAll().Return([]*models.User{
		{
			DN:          "CN=bob,DC=example,DC=nz",
			ID:          1,
			Username:    "bobe",
			DisplayName: "bobextample",
			Surname:     "Extample",
			GivenName:   "Bob",
			UIDNumber:   1001,
		},
	}, nil)

	suite.mockUserTable.EXPECT().Delete(models.User{
		ID:          1,
		DN:          "CN=bob,DC=example,DC=nz",
		Username:    "bobe",
		DisplayName: "bobextample",
		Surname:     "Extample",
		GivenName:   "Bob",
		UIDNumber:   1001,
	}).Once().Return(errors.New("failed"))

	response := api_handlers.SyncLDAPUsers(suite.mockDatabase, suite.mockLdap)
	require.NotNil(suite.T(), response)
	require.Equal(suite.T(), noodle_api.NewGetNoodleLdapReloadConflict().WithPayload(&models.Error{Message: "failed"}), response)
}

func (suite *LdapHandlersUserSyncTestSuite) TestHandlerSyncUsers_DBInsertError() {
	suite.mockDatabase.EXPECT().Tables().Return(suite.mockTables).Times(3)
	suite.mockTables.EXPECT().UserTable().Return(suite.mockUserTable).Times(3)

	suite.mockLdap.EXPECT().GetUsers().Return([]models.User{
		{
			DN:          "CN=bob,DC=example,DC=nz",
			Username:    "bobe",
			DisplayName: "bobextample",
			Surname:     "Extample",
			GivenName:   "Bob",
			UIDNumber:   1001,
		},
	}, nil)

	suite.mockUserTable.EXPECT().GetAll().Return([]*models.User{}, nil)

	suite.mockUserTable.EXPECT().ExistsDN("CN=bob,DC=example,DC=nz").Once().Return(false, nil)

	suite.mockUserTable.EXPECT().Insert(&models.User{
		DN:          "CN=bob,DC=example,DC=nz",
		Username:    "bobe",
		DisplayName: "bobextample",
		Surname:     "Extample",
		GivenName:   "Bob",
		UIDNumber:   1001,
	}).Once().Return(errors.New("failed"))

	response := api_handlers.SyncLDAPUsers(suite.mockDatabase, suite.mockLdap)
	require.NotNil(suite.T(), response)
	require.Equal(suite.T(), noodle_api.NewGetNoodleLdapReloadConflict().WithPayload(&models.Error{Message: "failed"}), response)
}

func TestLdapHandlersUserSyncTestSuite(t *testing.T) {
	suite.Run(t, new(LdapHandlersUserSyncTestSuite))
}
