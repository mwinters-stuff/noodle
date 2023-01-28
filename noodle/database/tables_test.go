package database_test

import (
	"errors"
	"testing"

	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/database/mocks"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TablesTestSuite struct {
	suite.Suite

	mockDatabase               *mocks.Database
	mockAppTemplateTable       *mocks.AppTemplateTable
	mockApplicationsTable      *mocks.ApplicationsTable
	mockApplicationTabTable    *mocks.ApplicationTabTable
	mockGroupApplicationsTable *mocks.GroupApplicationsTable
	mockGroupTable             *mocks.GroupTable
	mockTabTable               *mocks.TabTable
	mockUserApplicationsTable  *mocks.UserApplicationsTable
	mockUserGroupsTable        *mocks.UserGroupsTable
	mockUserTable              *mocks.UserTable
	mockUserSessionTable       *mocks.UserSessionTable
}

func (suite *TablesTestSuite) SetupSuite() {
	database.Logger = log.Output(nil)
}

func (suite *TablesTestSuite) SetupTest() {
	suite.mockDatabase = mocks.NewDatabase(suite.T())

	suite.mockAppTemplateTable = mocks.NewAppTemplateTable(suite.T())
	suite.mockApplicationsTable = mocks.NewApplicationsTable(suite.T())
	suite.mockApplicationTabTable = mocks.NewApplicationTabTable(suite.T())
	suite.mockGroupApplicationsTable = mocks.NewGroupApplicationsTable(suite.T())
	suite.mockGroupTable = mocks.NewGroupTable(suite.T())
	suite.mockTabTable = mocks.NewTabTable(suite.T())
	suite.mockUserApplicationsTable = mocks.NewUserApplicationsTable(suite.T())
	suite.mockUserGroupsTable = mocks.NewUserGroupsTable(suite.T())
	suite.mockUserTable = mocks.NewUserTable(suite.T())
	suite.mockUserSessionTable = mocks.NewUserSessionTable(suite.T())

	database.NewAppTemplateTable = func(database database.Database) database.AppTemplateTable { return suite.mockAppTemplateTable }

	database.NewApplicationsTable = func(database database.Database) database.ApplicationsTable { return suite.mockApplicationsTable }

	database.NewApplicationTabTable = func(database database.Database) database.ApplicationTabTable { return suite.mockApplicationTabTable }
	database.NewGroupApplicationsTable = func(database database.Database) database.GroupApplicationsTable {
		return suite.mockGroupApplicationsTable
	}
	database.NewGroupTable = func(database database.Database) database.GroupTable { return suite.mockGroupTable }
	database.NewTabTable = func(database database.Database) database.TabTable { return suite.mockTabTable }
	database.NewUserApplicationsTable = func(database database.Database) database.UserApplicationsTable {
		return suite.mockUserApplicationsTable
	}
	database.NewUserGroupsTable = func(database database.Database) database.UserGroupsTable { return suite.mockUserGroupsTable }
	database.NewUserTable = func(database database.Database) database.UserTable { return suite.mockUserTable }
	database.NewUserSessionTable = func(database database.Database) database.UserSessionTable { return suite.mockUserSessionTable }

}

func (suite *TablesTestSuite) TearDownTest() {
	database.NewAppTemplateTable = database.NewAppTemplateTableImpl
	database.NewApplicationsTable = database.NewApplicationsTableImpl
	database.NewApplicationTabTable = database.NewApplicationTabTableImpl
	database.NewGroupApplicationsTable = database.NewGroupApplicationsTableImpl
	database.NewGroupTable = database.NewGroupTableImpl
	database.NewTabTable = database.NewTabTableImpl
	database.NewUserApplicationsTable = database.NewUserApplicationsTableImpl
	database.NewUserGroupsTable = database.NewUserGroupsTableImpl
	database.NewUserTable = database.NewUserTableImpl
	database.NewUserSessionTable = database.NewUserSessionTableImpl

}

func (suite *TablesTestSuite) TestCreateTable() {

	suite.mockAppTemplateTable.EXPECT().Create().Once().Return(nil)
	suite.mockApplicationsTable.EXPECT().Create().Once().Return(nil)
	suite.mockApplicationTabTable.EXPECT().Create().Once().Return(nil)
	suite.mockGroupApplicationsTable.EXPECT().Create().Once().Return(nil)
	suite.mockGroupTable.EXPECT().Create().Once().Return(nil)
	suite.mockTabTable.EXPECT().Create().Once().Return(nil)
	suite.mockUserApplicationsTable.EXPECT().Create().Once().Return(nil)
	suite.mockUserGroupsTable.EXPECT().Create().Once().Return(nil)
	suite.mockUserTable.EXPECT().Create().Once().Return(nil)
	suite.mockUserSessionTable.EXPECT().Create().Once().Return(nil)

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Create()
	require.NoError(suite.T(), err)

	require.Equal(suite.T(), suite.mockAppTemplateTable, tables.AppTemplateTable())
	require.Equal(suite.T(), suite.mockApplicationsTable, tables.ApplicationsTable())
	require.Equal(suite.T(), suite.mockApplicationTabTable, tables.ApplicationTabTable())
	require.Equal(suite.T(), suite.mockGroupApplicationsTable, tables.GroupApplicationsTable())
	require.Equal(suite.T(), suite.mockGroupTable, tables.GroupTable())
	require.Equal(suite.T(), suite.mockTabTable, tables.TabTable())
	require.Equal(suite.T(), suite.mockUserApplicationsTable, tables.UserApplicationsTable())
	require.Equal(suite.T(), suite.mockUserGroupsTable, tables.UserGroupsTable())
	require.Equal(suite.T(), suite.mockUserTable, tables.UserTable())
	require.Equal(suite.T(), suite.mockUserSessionTable, tables.UserSessionTable())

}

func (suite *TablesTestSuite) TestCreateErrorAppTemplateTable() {
	suite.mockUserSessionTable.EXPECT().Create().Once().Return(nil)
	suite.mockAppTemplateTable.EXPECT().Create().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Create()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestCreateErrorApplicationsTable() {
	suite.mockUserSessionTable.EXPECT().Create().Once().Return(nil)
	suite.mockAppTemplateTable.EXPECT().Create().Once().Return(nil)
	suite.mockUserTable.EXPECT().Create().Once().Return(nil)
	suite.mockGroupTable.EXPECT().Create().Once().Return(nil)
	suite.mockApplicationsTable.EXPECT().Create().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Create()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestCreateErrorApplicationTabTable() {
	suite.mockUserSessionTable.EXPECT().Create().Once().Return(nil)
	suite.mockAppTemplateTable.EXPECT().Create().Once().Return(nil)
	suite.mockUserTable.EXPECT().Create().Once().Return(nil)
	suite.mockGroupTable.EXPECT().Create().Once().Return(nil)
	suite.mockApplicationsTable.EXPECT().Create().Once().Return(nil)
	suite.mockTabTable.EXPECT().Create().Once().Return(nil)
	suite.mockApplicationTabTable.EXPECT().Create().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Create()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestCreateErrorGroupApplicationsTable() {
	suite.mockUserSessionTable.EXPECT().Create().Once().Return(nil)
	suite.mockAppTemplateTable.EXPECT().Create().Once().Return(nil)
	suite.mockUserTable.EXPECT().Create().Once().Return(nil)
	suite.mockGroupTable.EXPECT().Create().Once().Return(nil)
	suite.mockApplicationsTable.EXPECT().Create().Once().Return(nil)
	suite.mockTabTable.EXPECT().Create().Once().Return(nil)
	suite.mockApplicationTabTable.EXPECT().Create().Once().Return(nil)
	suite.mockGroupApplicationsTable.EXPECT().Create().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Create()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestCreateErrorGroupTable() {
	suite.mockUserSessionTable.EXPECT().Create().Once().Return(nil)
	suite.mockAppTemplateTable.EXPECT().Create().Once().Return(nil)
	suite.mockUserTable.EXPECT().Create().Once().Return(nil)
	suite.mockGroupTable.EXPECT().Create().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Create()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestCreateErrorTabTable() {
	suite.mockUserSessionTable.EXPECT().Create().Once().Return(nil)
	suite.mockAppTemplateTable.EXPECT().Create().Once().Return(nil)
	suite.mockUserTable.EXPECT().Create().Once().Return(nil)
	suite.mockGroupTable.EXPECT().Create().Once().Return(nil)
	suite.mockApplicationsTable.EXPECT().Create().Once().Return(nil)
	suite.mockTabTable.EXPECT().Create().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Create()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestCreateErrorUserApplicationsTable() {
	suite.mockUserSessionTable.EXPECT().Create().Once().Return(nil)
	suite.mockAppTemplateTable.EXPECT().Create().Once().Return(nil)
	suite.mockUserTable.EXPECT().Create().Once().Return(nil)
	suite.mockGroupTable.EXPECT().Create().Once().Return(nil)
	suite.mockApplicationsTable.EXPECT().Create().Once().Return(nil)
	suite.mockTabTable.EXPECT().Create().Once().Return(nil)
	suite.mockApplicationTabTable.EXPECT().Create().Once().Return(nil)
	suite.mockGroupApplicationsTable.EXPECT().Create().Once().Return(nil)
	suite.mockUserApplicationsTable.EXPECT().Create().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Create()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestCreateErrorUserGroupsTable() {
	suite.mockUserSessionTable.EXPECT().Create().Once().Return(nil)
	suite.mockAppTemplateTable.EXPECT().Create().Once().Return(nil)
	suite.mockUserTable.EXPECT().Create().Once().Return(nil)
	suite.mockGroupTable.EXPECT().Create().Once().Return(nil)
	suite.mockApplicationsTable.EXPECT().Create().Once().Return(nil)
	suite.mockTabTable.EXPECT().Create().Once().Return(nil)
	suite.mockApplicationTabTable.EXPECT().Create().Once().Return(nil)
	suite.mockGroupApplicationsTable.EXPECT().Create().Once().Return(nil)
	suite.mockUserApplicationsTable.EXPECT().Create().Once().Return(nil)
	suite.mockUserGroupsTable.EXPECT().Create().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Create()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestCreateErrorUserTable() {
	suite.mockUserSessionTable.EXPECT().Create().Once().Return(nil)
	suite.mockAppTemplateTable.EXPECT().Create().Once().Return(nil)
	suite.mockUserTable.EXPECT().Create().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Create()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestCreateErrorUserSessionTable() {
	suite.mockUserSessionTable.EXPECT().Create().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Create()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestDropTable() {

	suite.mockApplicationTabTable.EXPECT().Drop().Once().Return(nil).NotBefore(suite.mockGroupApplicationsTable.ExpectedCalls...)
	suite.mockGroupApplicationsTable.EXPECT().Drop().Once().Return(nil).NotBefore(suite.mockUserApplicationsTable.ExpectedCalls...)
	suite.mockUserApplicationsTable.EXPECT().Drop().Once().Return(nil).NotBefore(suite.mockUserGroupsTable.ExpectedCalls...)
	suite.mockUserGroupsTable.EXPECT().Drop().Once().Return(nil).NotBefore(suite.mockUserTable.ExpectedCalls...)
	suite.mockUserTable.EXPECT().Drop().Once().Return(nil).NotBefore(suite.mockTabTable.ExpectedCalls...)
	suite.mockTabTable.EXPECT().Drop().Once().Return(nil).NotBefore(suite.mockGroupTable.ExpectedCalls...)
	suite.mockGroupTable.EXPECT().Drop().Once().Return(nil).NotBefore(suite.mockApplicationsTable.ExpectedCalls...)
	suite.mockApplicationsTable.EXPECT().Drop().Once().Return(nil).NotBefore(suite.mockAppTemplateTable.ExpectedCalls...)
	suite.mockAppTemplateTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserSessionTable.EXPECT().Drop().Once().Return(nil)

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Drop()
	require.NoError(suite.T(), err)
}

func (suite *TablesTestSuite) TestDropUserSessionTable() {
	suite.mockApplicationTabTable.EXPECT().Drop().Once().Return(nil)
	suite.mockGroupApplicationsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserApplicationsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserGroupsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserTable.EXPECT().Drop().Once().Return(nil)
	suite.mockTabTable.EXPECT().Drop().Once().Return(nil)
	suite.mockGroupTable.EXPECT().Drop().Once().Return(nil)
	suite.mockApplicationsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockAppTemplateTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserSessionTable.EXPECT().Drop().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Drop()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestDropErrorAppTemplateTable() {
	suite.mockApplicationTabTable.EXPECT().Drop().Once().Return(nil)
	suite.mockGroupApplicationsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserApplicationsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserGroupsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserTable.EXPECT().Drop().Once().Return(nil)
	suite.mockTabTable.EXPECT().Drop().Once().Return(nil)
	suite.mockGroupTable.EXPECT().Drop().Once().Return(nil)
	suite.mockApplicationsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockAppTemplateTable.EXPECT().Drop().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Drop()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestDropErrorApplicationsTable() {
	suite.mockApplicationTabTable.EXPECT().Drop().Once().Return(nil)
	suite.mockGroupApplicationsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserApplicationsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserGroupsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserTable.EXPECT().Drop().Once().Return(nil)
	suite.mockTabTable.EXPECT().Drop().Once().Return(nil)
	suite.mockGroupTable.EXPECT().Drop().Once().Return(nil)
	suite.mockApplicationsTable.EXPECT().Drop().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Drop()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestDropErrorUserTable() {
	suite.mockApplicationTabTable.EXPECT().Drop().Once().Return(nil)
	suite.mockGroupApplicationsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserApplicationsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserGroupsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserTable.EXPECT().Drop().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Drop()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestDropErrorApplicationTabTable() {
	suite.mockApplicationTabTable.EXPECT().Drop().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Drop()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestDropErrorGroupApplicationsTable() {
	suite.mockApplicationTabTable.EXPECT().Drop().Once().Return(nil)
	suite.mockGroupApplicationsTable.EXPECT().Drop().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Drop()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestDropErrorGroupTable() {
	suite.mockApplicationTabTable.EXPECT().Drop().Once().Return(nil)
	suite.mockGroupApplicationsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserApplicationsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserGroupsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserTable.EXPECT().Drop().Once().Return(nil)
	suite.mockTabTable.EXPECT().Drop().Once().Return(nil)
	suite.mockGroupTable.EXPECT().Drop().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Drop()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestDropErrorTabTable() {
	suite.mockApplicationTabTable.EXPECT().Drop().Once().Return(nil)
	suite.mockGroupApplicationsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserApplicationsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserGroupsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserTable.EXPECT().Drop().Once().Return(nil)
	suite.mockTabTable.EXPECT().Drop().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Drop()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestDropErrorUserApplicationsTable() {
	suite.mockApplicationTabTable.EXPECT().Drop().Once().Return(nil)
	suite.mockGroupApplicationsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserApplicationsTable.EXPECT().Drop().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Drop()
	require.EqualError(suite.T(), err, "failed")
}

func (suite *TablesTestSuite) TestDropErrorUserGroupsTable() {
	suite.mockApplicationTabTable.EXPECT().Drop().Once().Return(nil)
	suite.mockGroupApplicationsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserApplicationsTable.EXPECT().Drop().Once().Return(nil)
	suite.mockUserGroupsTable.EXPECT().Drop().Once().Return(errors.New("failed"))

	tables := database.NewTables()

	tables.InitTables(suite.mockDatabase)

	err := tables.Drop()
	require.EqualError(suite.T(), err, "failed")
}

func TestTablesSuite(t *testing.T) {
	suite.Run(t, new(TablesTestSuite))
}
