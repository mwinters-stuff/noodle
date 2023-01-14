package database_test

import (
	"net"
	"testing"

	"github.com/jackc/pgmock"
	database_test "github.com/mwinters-stuff/noodle/internal/database"
	"github.com/mwinters-stuff/noodle/server/models"

	"github.com/mwinters-stuff/noodle/noodle/database"
	"github.com/mwinters-stuff/noodle/noodle/yamltypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AppTemplateTableTestSuite struct {
	suite.Suite
	script        *pgmock.Script
	listener      net.Listener
	appConfig     yamltypes.AppConfig
	testFunctions database_test.TestFunctions
}

func (suite *AppTemplateTableTestSuite) SetupSuite() {
}

func (suite *AppTemplateTableTestSuite) SetupTest() {
	suite.testFunctions = database_test.TestFunctions{}
	suite.script = &pgmock.Script{
		Steps: pgmock.AcceptUnauthenticatedConnRequestSteps(),
	}

	suite.listener, suite.appConfig = suite.testFunctions.TestStepsRunner(suite.T(), suite.script)
}

func (suite *AppTemplateTableTestSuite) TearDownTest() {
	suite.listener.Close()
}

func (suite *AppTemplateTableTestSuite) TestCreateTable() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.CreateAppTemplateTableSteps(suite.T(), suite.script)

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewAppTemplateTable(db)

	err = table.Create()
	require.NoError(suite.T(), err)

}

func (suite *AppTemplateTableTestSuite) TestUpgrade() {
	table := database.NewAppTemplateTable(nil)
	require.Panics(suite.T(), func() { table.Upgrade(0, 0) })
}

func (suite *AppTemplateTableTestSuite) TestDrop() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Query","String":"DROP TABLE application_template"}`,
		`B {"Type":"CommandComplete","CommandTag":"DROP TABLE"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewAppTemplateTable(db)

	err = table.Drop()
	require.NoError(suite.T(), err)

}

func (suite *AppTemplateTableTestSuite) TestInsert() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_2","Query":"INSERT INTO application_template (appid,name,website,license,description,enhanced,tilebackground,icon,sha) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_2"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[1042,1043,1043,1043,1043,16,1043,1043,1042]}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_2","ParameterFormatCodes":[0,0,0,0,0,1,0,0,0],"Parameters":[{"text":"140902edbcc424c09736af28ab2de604c3bde936"},{"text":"AdGuard Home"},{"text":"https://github.com/AdguardTeam/AdGuardHome"},{"text":"GNU General Public License v3.0 only"},{"text":"AdGuard Home is a network-wide software for blocking ads."},{"binary":"01"},{"text":"light"},{"text":"adguardhome.png"},{"text":"ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7"}],"ResultFormatCodes":[]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"CommandComplete","CommandTag":"INSERT 0 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	app := models.ApplicationTemplate{
		Appid:          "140902edbcc424c09736af28ab2de604c3bde936",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
		SHA:            "ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7",
	}

	table := database.NewAppTemplateTable(db)

	err = table.Insert(app)
	require.NoError(suite.T(), err)

}

func (suite *AppTemplateTableTestSuite) TestUpdate() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"UPDATE application_template SET name = $2,website = $3,license = $4,description = $5,enhanced = $6,tilebackground = $7,icon = $8,sha = $9 WHERE appid = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[1042,1043,1043,1043,1043,16,1043,1043,1042]}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[0,0,0,0,0,1,0,0,0],"Parameters":[{"text":"140902edbcc424c09736af28ab2de604c3bde936"},{"text":"AdGuard Home"},{"text":"https://github.com/AdguardTeam/AdGuardHome"},{"text":"GNU General Public License v3.0 only"},{"text":"AdGuard Home is a network-wide software for blocking ads."},{"binary":"01"},{"text":"light"},{"text":"adguardhome.png"},{"text":"ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7"}],"ResultFormatCodes":[]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"CommandComplete","CommandTag":"UPDATE 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	app := models.ApplicationTemplate{
		Appid:          "140902edbcc424c09736af28ab2de604c3bde936",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
		SHA:            "ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7",
	}

	table := database.NewAppTemplateTable(db)

	err = table.Update(app)
	require.NoError(suite.T(), err)

}

func (suite *AppTemplateTableTestSuite) TestDelete() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)
	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"DELETE FROM application_template WHERE appid = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[1042]}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[0],"Parameters":[{"text":"140902edbcc424c09736af28ab2de604c3bde936"}],"ResultFormatCodes":[]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"NoData"}`,
		`B {"Type":"CommandComplete","CommandTag":"DELETE 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	app := models.ApplicationTemplate{
		Appid:          "140902edbcc424c09736af28ab2de604c3bde936",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
		SHA:            "ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7",
	}

	table := database.NewAppTemplateTable(db)

	err = table.Delete(app)
	require.NoError(suite.T(), err)

}

func (suite *AppTemplateTableTestSuite) TestSearch() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"SELECT * FROM application_template WHERE name LIKE $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[25]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"appid","TableOID":25146,"TableAttributeNumber":1,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0},{"Name":"name","TableOID":25146,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":25146,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":25146,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":25146,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":25146,"TableAttributeNumber":6,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":0},{"Name":"tilebackground","TableOID":25146,"TableAttributeNumber":7,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":25146,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"sha","TableOID":25146,"TableAttributeNumber":9,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[0],"Parameters":[{"text":"%Guard%"}],"ResultFormatCodes":[0,0,0,0,0,1,0,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"appid","TableOID":25146,"TableAttributeNumber":1,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0},{"Name":"name","TableOID":25146,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":25146,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":25146,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":25146,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":25146,"TableAttributeNumber":6,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":1},{"Name":"tilebackground","TableOID":25146,"TableAttributeNumber":7,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":25146,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"sha","TableOID":25146,"TableAttributeNumber":9,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"text":"140902edbcc424c09736af28ab2de604c3bde936"},{"text":"AdGuard Home"},{"text":"https://github.com/AdguardTeam/AdGuardHome"},{"text":"GNU General Public License v3.0 only"},{"text":"AdGuard Home is a network-wide software for blocking ads."},{"binary":"01"},{"text":"light"},{"text":"adguardhome.png"},{"text":"ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	app := models.ApplicationTemplate{
		Appid:          "140902edbcc424c09736af28ab2de604c3bde936",
		Name:           "AdGuard Home",
		Website:        "https://github.com/AdguardTeam/AdGuardHome",
		License:        "GNU General Public License v3.0 only",
		Description:    "AdGuard Home is a network-wide software for blocking ads.",
		Enhanced:       true,
		TileBackground: "light",
		Icon:           "adguardhome.png",
		SHA:            "ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7",
	}

	table := database.NewAppTemplateTable(db)

	result, err := table.Search("Guard")
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), result)

	require.ElementsMatch(suite.T(), []models.ApplicationTemplate{app}, result)

}

func (suite *AppTemplateTableTestSuite) TestSearchQueryFails() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"SELECT * FROM application_template WHERE name LIKE $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[25]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"appid","TableOID":25146,"TableAttributeNumber":1,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0},{"Name":"name","TableOID":25146,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":25146,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":25146,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":25146,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":25146,"TableAttributeNumber":6,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":0},{"Name":"tilebackground","TableOID":25146,"TableAttributeNumber":7,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":25146,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"sha","TableOID":25146,"TableAttributeNumber":9,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[0],"Parameters":[{"text":"%Guard%"}],"ResultFormatCodes":[0,0,0,0,0,1,0,0,0]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"appid","TableOID":25146,"TableAttributeNumber":1,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0},{"Name":"name","TableOID":25146,"TableAttributeNumber":2,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":24,"Format":0},{"Name":"website","TableOID":25146,"TableAttributeNumber":3,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"license","TableOID":25146,"TableAttributeNumber":4,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":104,"Format":0},{"Name":"description","TableOID":25146,"TableAttributeNumber":5,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":1004,"Format":0},{"Name":"enhanced","TableOID":25146,"TableAttributeNumber":6,"DataTypeOID":16,"DataTypeSize":1,"TypeModifier":-1,"Format":1},{"Name":"tilebackground","TableOID":25146,"TableAttributeNumber":7,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"icon","TableOID":25146,"TableAttributeNumber":8,"DataTypeOID":1043,"DataTypeSize":-1,"TypeModifier":260,"Format":0},{"Name":"sha","TableOID":25146,"TableAttributeNumber":9,"DataTypeOID":1042,"DataTypeSize":-1,"TypeModifier":44,"Format":0}]}`,
		`B {"Type":"DataRow","Values":[{"text":"140902edbcc424c09736af28ab2de604c3bde936"},{"text":"AdGuard Home"},{"text":"https://github.com/AdguardTeam/AdGuardHome"},{"text":"GNU General Public License v3.0 only"},{"text":"AdGuard Home is a network-wide software for blocking ads."},{"binary":"01"},{"text":"light"},{"text":"adguardhome.png"},{"text":"ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewAppTemplateTable(db)

	db.Close()

	result, err := table.Search(`A`)
	require.Error(suite.T(), err)
	require.Nil(suite.T(), result)

}

func (suite *AppTemplateTableTestSuite) TestExists() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"SELECT COUNT(*) FROM application_template WHERE appid = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[1042]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"count","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":20,"DataTypeSize":8,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[0],"Parameters":[{"text":"140902edbcc424c09736af28ab2de604c3bde936"}],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"count","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":20,"DataTypeSize":8,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"0000000000000001"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewAppTemplateTable(db)

	result, err := table.Exists("140902edbcc424c09736af28ab2de604c3bde936")
	require.NoError(suite.T(), err)
	require.True(suite.T(), result)

}

func (suite *AppTemplateTableTestSuite) TestNotExists() {
	suite.testFunctions.SetupConnectionSteps(suite.T(), suite.script)

	suite.testFunctions.LoadDatabaseSteps(suite.T(), suite.script, []string{
		`F {"Type":"Parse","Name":"stmtcache_3","Query":"SELECT COUNT(*) FROM application_template WHERE appid = $1","ParameterOIDs":null}`,
		`F {"Type":"Describe","ObjectType":"S","Name":"stmtcache_3"}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"ParseComplete"}`,
		`B {"Type":"ParameterDescription","ParameterOIDs":[1042]}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"count","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":20,"DataTypeSize":8,"TypeModifier":-1,"Format":0}]}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
		`F {"Type":"Bind","DestinationPortal":"","PreparedStatement":"stmtcache_3","ParameterFormatCodes":[0],"Parameters":[{"text":"140902edbcc424c09736af28ab2de604c3bde936"}],"ResultFormatCodes":[1]}`,
		`F {"Type":"Describe","ObjectType":"P","Name":""}`,
		`F {"Type":"Execute","Portal":"","MaxRows":0}`,
		`F {"Type":"Sync"}`,
		`B {"Type":"BindComplete"}`,
		`B {"Type":"RowDescription","Fields":[{"Name":"count","TableOID":0,"TableAttributeNumber":0,"DataTypeOID":20,"DataTypeSize":8,"TypeModifier":-1,"Format":1}]}`,
		`B {"Type":"DataRow","Values":[{"binary":"0000000000000000"}]}`,
		`B {"Type":"CommandComplete","CommandTag":"SELECT 1"}`,
		`B {"Type":"ReadyForQuery","TxStatus":"I"}`,
	})

	db := database.NewDatabase(suite.appConfig)
	assert.NotNil(suite.T(), db)
	defer db.Close()

	err := db.Connect()
	require.NoError(suite.T(), err)

	table := database.NewAppTemplateTable(db)

	result, err := table.Exists("140902edbcc424c09736af28ab2de604c3bde936")
	require.NoError(suite.T(), err)
	require.False(suite.T(), result)

}

func TestAppTemplateTableSuite(t *testing.T) {
	suite.Run(t, new(AppTemplateTableTestSuite))
}
