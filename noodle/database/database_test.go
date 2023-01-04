package database_test

import (
	"fmt"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/jackc/pgmock"
	"github.com/jackc/pgproto3/v2"
	dbf "github.com/mwinters-stuff/noodle/internal/database"
	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/yamltypes"
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
	loghook   databaseLogHook
	script    *pgmock.Script
	listener  net.Listener
	appConfig yamltypes.AppConfig
}

func (suite *DatabaseTestInitialSuite) SetupSuite() {
	suite.loghook = databaseLogHook{}
	database.Logger = log.Hook(&suite.loghook)

}

func (suite *DatabaseTestInitialSuite) SetupTest() {
	suite.script = &pgmock.Script{
		Steps: pgmock.AcceptUnauthenticatedConnRequestSteps(),
	}

	suite.listener, suite.appConfig = dbf.TestStepsRunner(suite.T(), suite.script)
}

func (suite *DatabaseTestInitialSuite) TearDownTest() {
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

	config, _ := yamltypes.UnmarshalConfig([]byte(yamltext))

	db := database.NewDatabase(config)
	assert.NotNil(suite.T(), db)

	dbf.SetupConnectionSteps(suite.script)
	err := db.Connect()
	require.Error(suite.T(), err)

	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.ErrorLevel && suite.loghook.LastMsg == "database connection failed failed to connect to `host=badhostname user=postgresuser database=postgres`: hostname resolving error (lookup badhostname: Temporary failure in name resolution)"
	}, time.Second*3, time.Millisecond*100)

}

func (suite *DatabaseTestInitialSuite) TestConnect() {

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()
	dbf.SetupConnectionSteps(suite.script)

	err := db.Connect()
	require.NoError(suite.T(), err)

	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.InfoLevel && suite.loghook.LastMsg == "database connected"
	}, time.Second, time.Millisecond*100)

}

func (suite *DatabaseTestInitialSuite) TestGetVersionMocked() {
	dbf.SetupConnectionSteps(suite.script)

	dbf.QueryMock(suite.script, "SELECT version FROM version",
		pgproto3.Bind{
			DestinationPortal:    "",
			PreparedStatement:    "stmtcache_?",
			ParameterFormatCodes: nil,
			Parameters:           nil,
			ResultFormatCodes:    []int16{1},
		},
		[]pgproto3.FieldDescription{
			{
				Name:                 []byte("version"),
				TableOID:             0,
				TableAttributeNumber: 1,
				DataTypeOID:          23,
				DataTypeSize:         4,
				TypeModifier:         -1,
				Format:               0,
			},
		},
		[][]byte{[]byte("1")})

	db := database.NewDatabase(suite.appConfig)
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
	dbf.SetupConnectionSteps(suite.script)

	dbf.QueryMock(suite.script, "SELECT version FROM version",
		pgproto3.Bind{
			DestinationPortal:    "",
			PreparedStatement:    "stmtcache_?",
			ParameterFormatCodes: nil,
			Parameters:           nil,
			ResultFormatCodes:    []int16{1},
		},
		[]pgproto3.FieldDescription{
			{
				Name:                 []byte("version"),
				TableOID:             0,
				TableAttributeNumber: 1,
				DataTypeOID:          23,
				DataTypeSize:         4,
				TypeModifier:         -1,
				Format:               0,
			},
		},
		[][]byte{[]byte(strconv.Itoa(database.DATABASE_VERSION))})

	db := database.NewDatabase(suite.appConfig)
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
	dbf.SetupConnectionSteps(suite.script)

	dbf.QueryMock(suite.script, "SELECT version FROM version",
		pgproto3.Bind{
			DestinationPortal:    "",
			PreparedStatement:    "stmtcache_?",
			ParameterFormatCodes: nil,
			Parameters:           nil,
			ResultFormatCodes:    []int16{1},
		},
		[]pgproto3.FieldDescription{
			{
				Name:                 []byte("version"),
				TableOID:             0,
				TableAttributeNumber: 1,
				DataTypeOID:          23,
				DataTypeSize:         4,
				TypeModifier:         -1,
				Format:               0,
			},
		},
		[][]byte{[]byte(strconv.Itoa(database.DATABASE_VERSION - 1))})

	db := database.NewDatabase(suite.appConfig)
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
	dbf.SetupConnectionSteps(suite.script)

	dbf.QueryMock(suite.script, "SELECT version FROM version",
		pgproto3.Bind{
			DestinationPortal:    "",
			PreparedStatement:    "stmtcache_?",
			ParameterFormatCodes: nil,
			Parameters:           nil,
			ResultFormatCodes:    []int16{1},
		},
		[]pgproto3.FieldDescription{
			{
				Name:                 []byte("version"),
				TableOID:             0,
				TableAttributeNumber: 1,
				DataTypeOID:          23,
				DataTypeSize:         4,
				TypeModifier:         -1,
				Format:               0,
			},
		},
		[][]byte{[]byte(strconv.Itoa(database.DATABASE_VERSION + 1))})

	db := database.NewDatabase(suite.appConfig)
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
	dbf.SetupConnectionSteps(suite.script)

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	err = db.Upgrade(database.DATABASE_VERSION - 1)
	require.NoError(suite.T(), err)
	assert.Eventually(suite.T(), func() bool {
		return suite.loghook.LastLevel == zerolog.InfoLevel && suite.loghook.LastMsg == fmt.Sprintf("upgrade database from %d to %d", database.DATABASE_VERSION-1, database.DATABASE_VERSION)
	}, time.Second, time.Millisecond*100)

}

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(DatabaseTestInitialSuite))
}
