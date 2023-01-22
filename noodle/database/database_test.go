package database_test

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/jackc/pgmock"
	database_test "github.com/mwinters-stuff/noodle/internal/database"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/database/mocks"
	"github.com/mwinters-stuff/noodle/noodle/options"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type databaseLogHook struct {
	LastEvent *zerolog.Event
	LastLevel zerolog.Level
	LastMsg   string
}

func (h *databaseLogHook) Run(e *zerolog.Event, l zerolog.Level, m string) {
	h.LastEvent = e
	h.LastLevel = l
	h.LastMsg = m
}

type DatabaseTestInitialSuite struct {
	suite.Suite
	loghook       databaseLogHook
	script        *pgmock.Script
	listener      net.Listener
	appConfig     options.AllNoodleOptions
	testFunctions database_test.TestFunctions
	tablesMock    *mocks.Tables
}

func (suite *DatabaseTestInitialSuite) SetupSuite() {

	suite.loghook = databaseLogHook{}
	database.Logger = log.Hook(&suite.loghook).Output(nil)

}

func (suite *DatabaseTestInitialSuite) SetupTest() {
	suite.testFunctions = database_test.TestFunctions{}
	suite.script = &pgmock.Script{
		Steps: pgmock.AcceptUnauthenticatedConnRequestSteps(),
	}

	suite.tablesMock = mocks.NewTables(suite.T())
	suite.listener, suite.appConfig = suite.testFunctions.TestStepsRunner(suite.T(), suite.script)
	database.NewTables = func() database.Tables { return suite.tablesMock }
}

func (suite *DatabaseTestInitialSuite) TearDownTest() {
	database.NewTables = database.NewTablesImpl
	suite.listener.Close()
}

func (suite *DatabaseTestInitialSuite) TearDownSuite() {

}

func (suite *DatabaseTestInitialSuite) TestBadConnect() {
	yamltext := `
postgres:
  user: postgresuser
  password: postgrespass
  db: postgres
  hostname: badhostname
  port: 1231
ldap:
  url: ldap://example.com
  base_dn: dc=example,dc=com
  username_attribute: uid
  additional_users_dn: ou=people
  users_filter: (&({username_attribute}={input})(objectClass=person))
  additional_groups_dn: ou=groups
  groups_filter: (&(uniquemember={dn})(objectclass=groupOfUniqueNames))
  group_name_attribute: cn
  display_name_attribute: displayName
  user: CN=readonly,DC=example,DC=com
  password: readonly
`

	config, _ := options.UnmarshalOptions([]byte(yamltext))

	db := database.NewDatabase(config.PostgresOptions)
	assert.NotNil(suite.T(), db)

	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	err := db.Connect()
	require.Error(suite.T(), err)

	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.ErrorLevel
	}, time.Second*3, time.Millisecond*100)

}

func (suite *DatabaseTestInitialSuite) TestConnect() {

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	err := db.Connect()
	require.NoError(suite.T(), err)

	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.InfoLevel && suite.loghook.LastMsg == "database connected"
	}, time.Second, time.Millisecond*100)

}

func (suite *DatabaseTestInitialSuite) TestCreated() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_2","Query":"SELECT version FROM version","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_2"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"version","TableOID":25129,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_2","ParameterFormatCodes":null,"Parameters":[],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"version","TableOID":25129,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	created, err := db.CheckCreated()

	require.NoError(suite.T(), err)
	assert.True(suite.T(), created)

}

func (suite *DatabaseTestInitialSuite) TestCreatedNotCreated() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_2","Query":"SELECT version FROM version","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_2"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ErrorResponse","Severity":"ERROR","SeverityUnlocalized":"ERROR","Code":"42P01","Message":"relation \"version\" does not exist","Detail":"","Hint":"","Position":21,"InternalPosition":0,"InternalQuery":"","Where":"","SchemaName":"","TableName":"","ColumnName":"","DataTypeName":"","ConstraintName":"","File":"parse_relation.c","Line":1392,"Routine":"parserOpenTable","UnknownFields":null}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	created, err := db.CheckCreated()

	require.Error(suite.T(), err)
	assert.False(suite.T(), created)

}

func (suite *DatabaseTestInitialSuite) TestGetVersion() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_2","Query":"SELECT version FROM version","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_2"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"version","TableOID":25129,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_2","ParameterFormatCodes":null,"Parameters":[],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"version","TableOID":25129,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"00000001"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	version, err := db.GetVersion()

	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, version)
	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.InfoLevel && suite.loghook.LastMsg == "current database version 1"
	}, time.Second, time.Millisecond*100)

}

func (suite *DatabaseTestInitialSuite) TestCheckUpgradeSameVersion() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_2","Query":"SELECT version FROM version","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_2"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"version","TableOID":25129,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_2","ParameterFormatCodes":null,"Parameters":[],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"version","TableOID":25129,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		fmt.Sprintf(`B {"Type":"DataRow","Values":[{"binary":"%08x"}]}`, database.DATABASE_VERSION),
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	upgrade, err := db.CheckUpgrade()
	require.NoError(suite.T(), err)
	require.False(suite.T(), upgrade)

	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.InfoLevel && suite.loghook.LastMsg == "no database upgrade required"
	}, time.Second, time.Millisecond*100)

}

func (suite *DatabaseTestInitialSuite) TestCheckUpgradeNewerVersion() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_2","Query":"SELECT version FROM version","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_2"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"version","TableOID":25129,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_2","ParameterFormatCodes":null,"Parameters":[],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"version","TableOID":25129,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		fmt.Sprintf(`B {"Type":"DataRow","Values":[{"binary":"%08x"}]}`, database.DATABASE_VERSION-1),
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	upgrade, err := db.CheckUpgrade()
	require.NoError(suite.T(), err)
	require.True(suite.T(), upgrade)

	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.InfoLevel && suite.loghook.LastMsg == fmt.Sprintf("upgrade database required from %d to %d", database.DATABASE_VERSION-1, database.DATABASE_VERSION)
	}, time.Second, time.Millisecond*100)

}

func (suite *DatabaseTestInitialSuite) TestCheckUpgradeDowngradeVersion() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_2","Query":"SELECT version FROM version","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_2"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"version","TableOID":25129,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_2","ParameterFormatCodes":null,"Parameters":[],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"version","TableOID":25129,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		fmt.Sprintf(`B {"Type":"DataRow","Values":[{"binary":"%08x"}]}`, database.DATABASE_VERSION+1),
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	upgrade, err := db.CheckUpgrade()
	require.ErrorContains(suite.T(), err, "datatabase downgrade not allowed")
	require.False(suite.T(), upgrade)

	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.ErrorLevel && suite.loghook.LastMsg == "cannot downgrade database"
	}, time.Second, time.Millisecond*100)

}

func (suite *DatabaseTestInitialSuite) TestUpgrade() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_2","Query":"SELECT version FROM version","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_2"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"version","TableOID":25129,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_2","ParameterFormatCodes":null,"Parameters":[],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"version","TableOID":25129,"TableAttributeNumber":1,"DataTypeOID":23,"DataTypeSize":4,"TypeModifier":-1,"Format":1}]}`,
		fmt.Sprintf(`B {"Type":"DataRow","Values":[{"binary":"%08x"}]}`, database.DATABASE_VERSION-1),
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	err = db.Upgrade()
	require.NoError(suite.T(), err)
	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.InfoLevel && suite.loghook.LastMsg == fmt.Sprintf("upgrade database from %d to %d", database.DATABASE_VERSION-1, database.DATABASE_VERSION)
	}, time.Second, time.Millisecond*100)

}

func (suite *DatabaseTestInitialSuite) TestCreate() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.tablesMock.EXPECT().Create().Once().Return(nil)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		fmt.Sprintf(`F {"Type":"Query","String":"\nCREATE TABLE IF NOT EXISTS version (version int);\nDELETE FROM version;\nINSERT INTO version (version) values (%d)\n"}`, database.DATABASE_VERSION),
		`B {"Type":"CommandComplete","CommandTag":"CREATE TABLE"}`,
		`B {"Type":"CommandComplete","CommandTag":"DELETE 0"}`,
		`B {"Type":"CommandComplete","CommandTag":"INSERT 0 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	err = db.Create()
	require.NoError(suite.T(), err)

}

func (suite *DatabaseTestInitialSuite) TestCreateError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		fmt.Sprintf(`F {"Type":"Query","String":"\nCREATE TABLE IF NOT EXISTS version (version int);\nDELETE FROM version;\nINSERT INTO version (version) values (%d)\n"}`, database.DATABASE_VERSION),
		`B {"Type":"ErrorResponse","Severity":"ERROR","SeverityUnlocalized":"ERROR","Code":"42P01","Message":"relation \"version\" does not exist","Detail":"","Hint":"","Position":21,"InternalPosition":0,"InternalQuery":"","Where":"","SchemaName":"","TableName":"","ColumnName":"","DataTypeName":"","ConstraintName":"","File":"parse_relation.c","Line":1392,"Routine":"parserOpenTable","UnknownFields":null}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	err = db.Create()
	require.Error(suite.T(), err)

}

func (suite *DatabaseTestInitialSuite) TestDrop() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.tablesMock.EXPECT().Drop().Once().Return(nil)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Query","String":"DROP TABLE version"}`,
		`B {"Type":"CommandComplete","CommandTag":"DROP TABLE"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	err = db.Drop()
	require.NoError(suite.T(), err)
}

func (suite *DatabaseTestInitialSuite) TestDropError() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Query","String":"DROP TABLE version"}`,
		`B {"Type":"ErrorResponse","Severity":"ERROR","SeverityUnlocalized":"ERROR","Code":"42P01","Message":"relation \"version\" does not exist","Detail":"","Hint":"","Position":21,"InternalPosition":0,"InternalQuery":"","Where":"","SchemaName":"","TableName":"","ColumnName":"","DataTypeName":"","ConstraintName":"","File":"parse_relation.c","Line":1392,"Routine":"parserOpenTable","UnknownFields":null}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	err = db.Drop()
	require.Error(suite.T(), err)
}

func (suite *DatabaseTestInitialSuite) TestTables() {
	db := database.NewDatabase(suite.appConfig.PostgresOptions)
	assert.NotNil(suite.T(), db)

	tables := database.NewTables()
	dbTables := db.Tables()

	require.IsType(suite.T(), tables, dbTables)
}

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(DatabaseTestInitialSuite))
}
